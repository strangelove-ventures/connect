package alpaca

import (
	"fmt"
	"math/big"
	"time"

	providertypes "github.com/skip-mev/connect/v2/providers/types"
	"go.uber.org/zap"

	"github.com/skip-mev/connect/v2/oracle/types"
)

// parseQuotes handles an array of quotes from Alpaca
func (h *WebSocketHandler) parseQuotes(quotes []Quote) (types.PriceResponse, error) {
	var (
		resolved   = make(types.ResolvedPrices)
		unResolved = make(types.UnResolvedPrices)
	)

	// Process each quote
	for _, quote := range quotes {
		// Determine if the ticker is valid.
		ticker, ok := h.cache.FromOffChainTicker(quote.Symbol)
		if !ok {
			h.logger.Debug("got response for an unsupported symbol", zap.String("symbol", quote.Symbol))
			continue
		}

		// Calculate mid price (average of bid and ask)
		if quote.BidPrice <= 0 || quote.AskPrice <= 0 {
			unResolved[ticker] = providertypes.UnresolvedResult{
				ErrorWithCode: providertypes.NewErrorWithCode(
					fmt.Errorf("invalid quote prices: bid=%f, ask=%f", quote.BidPrice, quote.AskPrice),
					providertypes.ErrorFailedToParsePrice),
			}
			continue
		}

		midPrice := (quote.BidPrice + quote.AskPrice) / 2
		price := big.NewFloat(midPrice)

		// TODO: USE RESPONSE'S TIMESTAMP WHEN WE GET A NEW API TOKEN
		// WHICH CAN HANDLE REAL TIME PRICES (sip)
		//
		// // Convert timestamp to time (Alpaca uses RFC3339 format)
		// timestamp, err := time.Parse(time.RFC3339, quote.Timestamp)
		// if err != nil {
		// 	unResolved[ticker] = providertypes.UnresolvedResult{
		// 		ErrorWithCode: providertypes.NewErrorWithCode(
		// 			fmt.Errorf("invalid timestamp format: %s", quote.Timestamp),
		// 			providertypes.ErrorFailedToDecode),
		// 	}
		// 	continue
		// }
		//
		// // Add to resolved prices
		// resolved[ticker] = types.NewPriceResult(price, timestamp)

		// Add to resolved prices
		resolved[ticker] = types.NewPriceResult(price, time.Now().UTC())

		// Debug the specific ticker, price, and timestamp
		h.logger.Debug("Successfully processed Alpaca quote",
			zap.String("ticker", ticker.String()),
			zap.String("offChainTicker", ticker.GetOffChainTicker()),
			zap.String("price", price.String()))
	}

	return types.NewPriceResponse(resolved, unResolved), nil
}
