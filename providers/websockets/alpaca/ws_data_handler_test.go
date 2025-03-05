package alpaca_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/skip-mev/connect/v2/oracle/config"
	"github.com/skip-mev/connect/v2/oracle/types"
	"github.com/skip-mev/connect/v2/providers/base/websocket/handlers"
	providertypes "github.com/skip-mev/connect/v2/providers/types"
	"github.com/skip-mev/connect/v2/providers/websockets/alpaca"
)

var (
	aaplUsd = types.DefaultProviderTicker{
		OffChainTicker: "AAPL",
	}
	googUsd = types.DefaultProviderTicker{
		OffChainTicker: "GOOGL",
	}
	unknown = types.DefaultProviderTicker{
		OffChainTicker: "UNKNOWN",
	}
	logger = zap.NewExample()
)

func TestHandleMessage(t *testing.T) {
	testCases := []struct {
		name          string
		msg           func() []byte
		resp          types.PriceResponse
		updateMessage func() []handlers.WebsocketEncodedMessage
		expErr        bool
	}{
		{
			name: "authentication success response",
			msg: func() []byte {
				msg := []alpaca.AuthenticatedMessage{
					{
						Type: "success",
						Msg:  "authenticated",
					},
				}

				bz, err := json.Marshal(msg)
				require.NoError(t, err)

				return bz
			},
			resp:          types.PriceResponse{},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false,
		},
		{
			name: "subscription confirmation response",
			msg: func() []byte {
				msg := []alpaca.SubscriptionConfirmationMessage{
					{
						Type:   "subscription",
						Quotes: []string{"AAPL", "GOOGL"},
					},
				}

				bz, err := json.Marshal(msg)
				require.NoError(t, err)

				return bz
			},
			resp:          types.PriceResponse{},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false,
		},
		{
			name: "quote update with valid prices",
			msg: func() []byte {
				msg := []alpaca.Quote{
					{
						Type:        "q",
						Symbol:      "AAPL",
						BidExchange: "N",
						BidPrice:    150.25,
						BidSize:     100,
						AskExchange: "N",
						AskPrice:    150.35,
						AskSize:     100,
						Conditions:  []string{},
						Tape:        "A",
						Timestamp:   "2023-04-20T15:30:45.123456789Z",
					},
				}

				bz, err := json.Marshal(msg)
				require.NoError(t, err)

				return bz
			},
			resp: types.PriceResponse{
				Resolved: types.ResolvedPrices{
					aaplUsd: {
						Value: big.NewFloat(150.3), // Average of bid and ask
					},
				},
			},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false,
		},
		{
			name: "quote update with multiple symbols",
			msg: func() []byte {
				msg := []alpaca.Quote{
					{
						Type:        "q",
						Symbol:      "AAPL",
						BidExchange: "N",
						BidPrice:    150.25,
						BidSize:     100,
						AskExchange: "N",
						AskPrice:    150.35,
						AskSize:     100,
						Conditions:  []string{},
						Tape:        "A",
						Timestamp:   "2023-04-20T15:30:45.123456789Z",
					},
					{
						Type:        "q",
						Symbol:      "GOOGL",
						BidExchange: "N",
						BidPrice:    2500.50,
						BidSize:     10,
						AskExchange: "N",
						AskPrice:    2501.50,
						AskSize:     10,
						Conditions:  []string{},
						Tape:        "A",
						Timestamp:   "2023-04-20T15:30:45.123456789Z",
					},
				}

				bz, err := json.Marshal(msg)
				require.NoError(t, err)

				return bz
			},
			resp: types.PriceResponse{
				Resolved: types.ResolvedPrices{
					aaplUsd: {
						Value: big.NewFloat(150.3), // Average of bid and ask
					},
					googUsd: {
						Value: big.NewFloat(2501.0), // Average of bid and ask
					},
				},
			},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false,
		},
		{
			name: "quote update with unknown symbol",
			msg: func() []byte {
				msg := []alpaca.Quote{
					{
						Type:        "q",
						Symbol:      "UNKNOWN",
						BidExchange: "N",
						BidPrice:    100.25,
						BidSize:     100,
						AskExchange: "N",
						AskPrice:    100.35,
						AskSize:     100,
						Conditions:  []string{},
						Tape:        "A",
						Timestamp:   "2023-04-20T15:30:45.123456789Z",
					},
				}

				bz, err := json.Marshal(msg)
				require.NoError(t, err)

				return bz
			},
			resp:          types.PriceResponse{},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false, // This shouldn't error, it will just skip unknown symbols
		},
		{
			name: "quote update with invalid prices",
			msg: func() []byte {
				msg := []alpaca.Quote{
					{
						Type:        "q",
						Symbol:      "AAPL",
						BidExchange: "N",
						BidPrice:    -10.25, // Negative bid, invalid
						BidSize:     100,
						AskExchange: "N",
						AskPrice:    150.35,
						AskSize:     100,
						Conditions:  []string{},
						Tape:        "A",
						Timestamp:   "2023-04-20T15:30:45.123456789Z",
					},
				}

				bz, err := json.Marshal(msg)
				require.NoError(t, err)

				return bz
			},
			resp: types.PriceResponse{
				UnResolved: types.UnResolvedPrices{
					aaplUsd: providertypes.UnresolvedResult{
						ErrorWithCode: providertypes.NewErrorWithCode(
							fmt.Errorf("invalid quote prices: bid=%f, ask=%f", -10.25, 150.35),
							providertypes.ErrorFailedToParsePrice),
					},
				},
			},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false,
		},
		{
			name: "quote update with zero ask price",
			msg: func() []byte {
				msg := []alpaca.Quote{
					{
						Type:        "q",
						Symbol:      "AAPL",
						BidExchange: "N",
						BidPrice:    150.25,
						BidSize:     100,
						AskExchange: "N",
						AskPrice:    0, // Zero ask, invalid
						AskSize:     100,
						Conditions:  []string{},
						Tape:        "A",
						Timestamp:   "2023-04-20T15:30:45.123456789Z",
					},
				}

				bz, err := json.Marshal(msg)
				require.NoError(t, err)

				return bz
			},
			resp: types.PriceResponse{
				UnResolved: types.UnResolvedPrices{
					aaplUsd: providertypes.UnresolvedResult{
						ErrorWithCode: providertypes.NewErrorWithCode(
							fmt.Errorf("invalid quote prices: bid=%f, ask=%f", 150.25, 0.0),
							providertypes.ErrorFailedToParsePrice),
					},
				},
			},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false,
		},
		{
			name: "unknown message format",
			msg: func() []byte {
				return []byte(`{"type":"unknown"}`)
			},
			resp:          types.PriceResponse{},
			updateMessage: func() []handlers.WebsocketEncodedMessage { return nil },
			expErr:        false, // Alpaca handler doesn't error on unknown messages
		},
	}

	wsHandler, err := alpaca.NewWebSocketDataHandler(logger, alpaca.DefaultWebSocketConfig)
	require.NoError(t, err)

	// Update the cache since it is assumed that CreateMessages is executed before anything else.
	_, err = wsHandler.CreateMessages([]types.ProviderTicker{aaplUsd, googUsd})
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, updateMsg, err := wsHandler.HandleMessage(tc.msg())
			if tc.expErr {
				require.Error(t, err)
				return
			}
			
			require.NoError(t, err)
			require.Equal(t, tc.updateMessage(), updateMsg)

			require.Equal(t, len(tc.resp.Resolved), len(resp.Resolved))
			require.Equal(t, len(tc.resp.UnResolved), len(resp.UnResolved))

			for cp, result := range tc.resp.Resolved {
				require.Contains(t, resp.Resolved, cp)
				
				// We can't directly compare timestamps since they are generated at runtime
				// Just verify the price values match
				require.Equal(t, result.Value.SetPrec(18), resp.Resolved[cp].Value.SetPrec(18))
				
				// Note: we're not comparing timestamps as they're set to time.Now() in the implementation
			}

			for cp, expectedErr := range tc.resp.UnResolved {
				require.Contains(t, resp.UnResolved, cp)
				require.Error(t, resp.UnResolved[cp])
				require.Equal(t, expectedErr.Error(), resp.UnResolved[cp].Error())
			}
		})
	}
}

func TestCreateMessages(t *testing.T) {
	batchCfg := alpaca.DefaultWebSocketConfig
	batchCfg.MaxSubscriptionsPerBatch = 2

	testCases := []struct {
		name        string
		cps         []types.ProviderTicker
		cfg         config.WebSocketConfig
		expectedLen int // Check the expected number of messages
		expectedErr bool
	}{
		{
			name:        "no currency pairs to subscribe to",
			cps:         []types.ProviderTicker{},
			cfg:         alpaca.DefaultWebSocketConfig,
			expectedLen: 0,
			expectedErr: true,
		},
		{
			name: "one currency pair to subscribe to",
			cps: []types.ProviderTicker{
				aaplUsd,
			},
			cfg:         alpaca.DefaultWebSocketConfig,
			expectedLen: 2, // Auth message + 1 subscription message
			expectedErr: false,
		},
		{
			name: "multiple currency pairs to subscribe to",
			cps: []types.ProviderTicker{
				aaplUsd,
				googUsd,
			},
			cfg:         alpaca.DefaultWebSocketConfig,
			expectedLen: 3, // Auth message + 2 subscription messages (one per ticker)
			expectedErr: false,
		},
		{
			name: "multiple currency pairs with batch config",
			cps: []types.ProviderTicker{
				aaplUsd,
				googUsd,
			},
			cfg:         batchCfg,
			expectedLen: 2, // Auth message + 1 batched subscription message
			expectedErr: false,
		},
		{
			name: "multiple currency pairs with batch config + 1",
			cps: []types.ProviderTicker{
				aaplUsd,
				googUsd,
				unknown,
			},
			cfg:         batchCfg,
			expectedLen: 3, // Auth message + 2 batched subscription messages
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wsHandler, err := alpaca.NewWebSocketDataHandler(logger, tc.cfg)
			require.NoError(t, err)

			msgs, err := wsHandler.CreateMessages(tc.cps)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedLen, len(msgs))
			
			// Verify first message is auth message
			if tc.expectedLen > 0 {
				var authMsg alpaca.AuthMessage
				err = json.Unmarshal(msgs[0], &authMsg)
				require.NoError(t, err)
				require.Equal(t, alpaca.AuthAction, authMsg.Action)
			}

			// Verify subscription messages
			for i := 1; i < len(msgs); i++ {
				var subMsg alpaca.SubscribeMessage
				err = json.Unmarshal(msgs[i], &subMsg)
				require.NoError(t, err)
				require.Equal(t, alpaca.SubscribeAction, subMsg.Action)
				require.NotEmpty(t, subMsg.Quotes)
				
				// If we're using batch config, check that batch size is respected
				if tc.cfg.MaxSubscriptionsPerBatch < len(tc.cps) {
					require.LessOrEqual(t, len(subMsg.Quotes), tc.cfg.MaxSubscriptionsPerBatch)
				}
			}
		})
	}
}
