package polygon

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/skip-mev/connect/v2/oracle/config"
	"github.com/skip-mev/connect/v2/oracle/types"
	providertypes "github.com/skip-mev/connect/v2/providers/types"
	"go.uber.org/zap"
)

var _ types.PriceAPIDataHandler = (*APIHandler)(nil)

// APIHandler implements the PriceAPIDataHandler interface for Polygon, which can be used
// by a base provider. The DataHandler fetches data from the Quote Latest V2 Polygon API.
// Requests for prices are fulfilled in a single request.
type APIHandler struct {
	// api is the config for the Polygon API.
	api config.APIConfig
	// cache maintains the latest set of tickers seen by the handler.
	cache types.ProviderTickers
}

// NewAPIHandler returns a new Polygon PriceAPIDataHandler.
func NewAPIHandler(
	logger *zap.Logger,
	api config.APIConfig,
) (types.PriceAPIDataHandler, error) {
	if api.Name != Name {
		return nil, fmt.Errorf("expected api config name %s, got %s", Name, api.Name)
	}

	if !api.Enabled {
		return nil, fmt.Errorf("api config for %s is not enabled", Name)
	}

	if err := api.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid api config for %s: %w", Name, err)
	}

	apiKey := os.Getenv("POLYGON_API_KEY")
	if apiKey == "" {
		logger.Error("missing Polygon API key", zap.String("POLYGON_API_KEY", apiKey))
	}

	// Set the API key for the Polygon API.
	if len(api.Endpoints) == 0 {
		api.Endpoints[0].Authentication = config.Authentication{
			APIKey:       fmt.Sprintf("Bearer %s", apiKey),
			APIKeyHeader: "Authorization",
		}
	}

	return &APIHandler{
		api:   api,
		cache: types.NewProviderTickers(),
	}, nil
}

// CreateURL returns the URL that is used to fetch data from the Polygon API for the
// given tickers.
func (h *APIHandler) CreateURL(
	tickers []types.ProviderTicker,
) (string, error) {
	var ids []string //nolint:prealloc
	for _, ticker := range tickers {
		ids = append(ids, ticker.GetOffChainTicker())
		h.cache.Add(ticker)
	}

	if len(ids) == 0 {
		return "", fmt.Errorf("no tickers provided")
	}

	if len(ids) > 1 {
		return "", fmt.Errorf("polygon.io only supports one ticker per request")
	}

	ticker := strings.Join(ids, ",")

	// Date must be formatted as "YYYY-MM-DD"
	// Date must be exactly 1 day ago from now.
	dateStr := time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02")

	url := fmt.Sprintf(Endpoint, h.api.Endpoints[0].URL, ticker, dateStr, dateStr)
	return url, nil
}

// ParseResponse parses the spot price HTTP response from the Polygon API and returns
// the resulting price(s).
func (h *APIHandler) ParseResponse(
	tickers []types.ProviderTicker,
	resp *http.Response,
) types.PriceResponse {
	// Parse the response into a CoinBaseResponse.
	var result PolygonResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.NewPriceResponseWithErr(
			tickers,
			providertypes.NewErrorWithCode(err, providertypes.ErrorFailedToDecode),
		)
	}

	if result.Status != "OK" {
		return types.NewPriceResponseWithErr(
			tickers,
			providertypes.NewErrorWithCode(
				fmt.Errorf("polygon error: %s", result.Status),
				providertypes.ErrorAPIGeneral,
			),
		)
	}

	var (
		resolved   = make(types.ResolvedPrices)
		unresolved = make(types.UnResolvedPrices)
	)

	// Get open quote for POC
	for _, tickerResponse := range result.Results {
		ticker, exists := h.cache.FromOffChainTicker(result.Ticker)
		if !exists {
			continue
		}

		quote := (tickerResponse.Open + tickerResponse.Close + tickerResponse.Low + tickerResponse.High) / 4
		resolved[ticker] = types.NewPriceResult(big.NewFloat(quote), time.Now().UTC())
	}

	return types.NewPriceResponse(resolved, unresolved)
}
