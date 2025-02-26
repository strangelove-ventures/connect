package polygon_test

import (
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	providertypes "github.com/skip-mev/connect/v2/providers/types"

	"github.com/stretchr/testify/require"

	"github.com/skip-mev/connect/v2/oracle/config"
	"github.com/skip-mev/connect/v2/oracle/types"
	"github.com/skip-mev/connect/v2/providers/apis/polygon"
	"github.com/skip-mev/connect/v2/providers/base/testutils"
)

var (
	amzxusd = types.DefaultProviderTicker{
		OffChainTicker: "AMZN",
	}
	nvdaxusd = types.DefaultProviderTicker{
		OffChainTicker: "NVDA",
	}
)

func TestCreateURL(t *testing.T) {
	testCases := []struct {
		name        string
		cps         []types.ProviderTicker
		url         string
		expectedErr bool
	}{
		{
			name:        "empty",
			cps:         []types.ProviderTicker{},
			url:         "",
			expectedErr: true,
		},
		{
			name: "valid single ticker",
			cps: []types.ProviderTicker{
				amzxusd,
			},
			url:         fmt.Sprintf("https://api.polygon.io/v2/aggs/ticker/AMZN/range/1/day/2024-01-09/2024-01-09"),
			expectedErr: false,
		},
		{
			name: "multiple tickers",
			cps: []types.ProviderTicker{
				amzxusd,
				nvdaxusd,
			},
			url:         fmt.Sprintf("https://api.polygon.io/v2/aggs/ticker/AMZN,NVDA/range/1/day/2024-01-09/2024-01-09"),
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h, err := polygon.NewAPIHandler(polygon.DefaultAPIConfig)
			require.NoError(t, err)

			url, err := h.CreateURL(tc.cps)
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.url, url)
			}
		})
	}
}

func TestParseResponse(t *testing.T) {
	testCases := []struct {
		name     string
		cps      []types.ProviderTicker
		response *http.Response
		expected types.PriceResponse
	}{
		{
			name: "valid response for single ticker",
			cps: []types.ProviderTicker{
				amzxusd,
			},
			response: testutils.CreateResponseFromJSON(
				`
{
    "ticker": "AMZN",
    "queryCount": 1,
    "resultsCount": 1,
    "adjusted": true,
    "results": [
        {
            "v": 28578710,
            "vw": 178.2351,
            "o": 175.2,
            "c": 179.5,
            "h": 179.99,
            "l": 174.28,
            "t": 1704844800000,
            "n": 332126
        }
    ],
    "status": "OK",
    "request_id": "6af5c3dbe10f50f9e4d38dbc4e2ef47b",
    "count": 1
}
				`,
			),
			expected: types.NewPriceResponse(
				types.ResolvedPrices{
					amzxusd: {
						Value: big.NewFloat(175.2), // Using open price
					},
				},
				types.UnResolvedPrices{},
			),
		},
		{
			name: "error status response",
			cps: []types.ProviderTicker{
				amzxusd,
			},
			response: testutils.CreateResponseFromJSON(
				`
{
    "ticker": "AMZN",
    "status": "ERROR",
    "request_id": "6af5c3dbe10f50f9e4d38dbc4e2ef47b",
    "count": 0
}
				`,
			),
			expected: types.NewPriceResponse(
				types.ResolvedPrices{},
				types.UnResolvedPrices{
					amzxusd: providertypes.UnresolvedResult{
						ErrorWithCode: providertypes.NewErrorWithCode(fmt.Errorf("polygon error: ERROR"), providertypes.ErrorAPIGeneral),
					},
				},
			),
		},
		{
			name: "malformed json response",
			cps: []types.ProviderTicker{
				amzxusd,
			},
			response: testutils.CreateResponseFromJSON(
				`
{
    "ticker": "AMZN",
    "status": "OK",
    "results": [
        {
            "v": 28578710,
            "vw": 178.2351,
            "o": 175.2,
            "c": 179.5,
            "h": 179.99,
            "l": 174.28,
            "t": 1704844800000,
            "n": 332126
        },
    ],
}
				`,
			),
			expected: types.NewPriceResponse(
				types.ResolvedPrices{},
				types.UnResolvedPrices{
					amzxusd: providertypes.UnresolvedResult{
						ErrorWithCode: providertypes.NewErrorWithCode(fmt.Errorf("bad format"), providertypes.ErrorFailedToDecode),
					},
				},
			),
		},
		{
			name: "unable to parse json",
			cps: []types.ProviderTicker{
				amzxusd,
			},
			response: testutils.CreateResponseFromJSON(
				`not a valid json response`,
			),
			expected: types.NewPriceResponse(
				types.ResolvedPrices{},
				types.UnResolvedPrices{
					amzxusd: providertypes.UnresolvedResult{
						ErrorWithCode: providertypes.NewErrorWithCode(fmt.Errorf("bad format"), providertypes.ErrorFailedToDecode),
					},
				},
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h, err := polygon.NewAPIHandler(polygon.DefaultAPIConfig)
			require.NoError(t, err)

			// First register the tickers with the handler's cache
			_, err = h.CreateURL(tc.cps)
			require.NoError(t, err)

			now := time.Now()
			resp := h.ParseResponse(tc.cps, tc.response)

			require.Len(t, resp.Resolved, len(tc.expected.Resolved))
			require.Len(t, resp.UnResolved, len(tc.expected.UnResolved))

			for cp, result := range tc.expected.Resolved {
				require.Contains(t, resp.Resolved, cp)
				r := resp.Resolved[cp]
				require.Equal(t, result.Value.SetPrec(18), r.Value.SetPrec(18))
				require.True(t, r.Timestamp.After(now) || r.Timestamp.Equal(now))
			}

			for cp := range tc.expected.UnResolved {
				require.Contains(t, resp.UnResolved, cp)
				require.Error(t, resp.UnResolved[cp])
			}
		})
	}
}

func TestNewAPIHandler(t *testing.T) {
	testCases := []struct {
		name        string
		config      config.APIConfig
		expectedErr bool
	}{
		{
			name:        "valid config",
			config:      polygon.DefaultAPIConfig,
			expectedErr: false,
		},
		{
			name: "incorrect name",
			config: func() config.APIConfig {
				cfg := polygon.DefaultAPIConfig
				cfg.Name = "wrong_name"
				return cfg
			}(),
			expectedErr: true,
		},
		{
			name: "disabled config",
			config: func() config.APIConfig {
				cfg := polygon.DefaultAPIConfig
				cfg.Enabled = false
				return cfg
			}(),
			expectedErr: true,
		},
		{
			name: "invalid endpoint",
			config: func() config.APIConfig {
				cfg := polygon.DefaultAPIConfig
				cfg.Endpoints = []config.Endpoint{}
				return cfg
			}(),
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := polygon.NewAPIHandler(tc.config)
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
