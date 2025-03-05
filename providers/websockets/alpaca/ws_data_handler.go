package alpaca

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/skip-mev/connect/v2/oracle/config"
	"github.com/skip-mev/connect/v2/oracle/types"
	"github.com/skip-mev/connect/v2/providers/base/websocket/handlers"
)

var _ types.PriceWebSocketDataHandler = (*WebSocketHandler)(nil)

// WebSocketHandler implements the WebSocketDataHandler interface. This is used to
// handle messages received from the Alpaca websocket API.
type WebSocketHandler struct {
	logger *zap.Logger

	// ws is the config for the Alpaca websocket.
	ws config.WebSocketConfig

	// apiKey and apiSecret for Alpaca authentication
	apiKey    string
	apiSecret string

	// cache maintains the latest set of tickers seen by the handler.
	cache types.ProviderTickers
}

// NewWebSocketDataHandler returns a new Alpaca PriceWebSocketDataHandler.
func NewWebSocketDataHandler(
	logger *zap.Logger,
	ws config.WebSocketConfig,
) (types.PriceWebSocketDataHandler, error) {
	if ws.Name != Name {
		return nil, fmt.Errorf("expected websocket config name %s, got %s", Name, ws.Name)
	}

	if !ws.Enabled {
		return nil, fmt.Errorf("websocket config for %s is not enabled", Name)
	}

	if err := ws.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid websocket config for %s: %w", Name, err)
	}

	apiKey := os.Getenv("ALPACA_API_KEY")
	apiSecret := os.Getenv("ALPACA_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		logger.Error("missing Alpaca API key and secret", zap.String("ALPACA_API_KEY", apiKey), zap.String("ALPACA_API_SECRET", apiSecret))
	}

	logger.Info("creating new Alpaca websocket handler",
		zap.String("endpoint", ws.Endpoints[0].URL),
		zap.Duration("read_timeout", ws.ReadTimeout),
		zap.Duration("reconnection_timeout", ws.ReconnectionTimeout))

	return &WebSocketHandler{
		logger:    logger,
		ws:        ws,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		cache:     types.NewProviderTickers(),
	}, nil
}

// HandleMessage is used to handle a message received from the Alpaca websocket API.
// The handler processes the following message flows:
// 1. Auth response after sending authentication credentials
// 2. Quote updates after subscribing to quotes
func (h *WebSocketHandler) HandleMessage(
	message []byte,
) (types.PriceResponse, []handlers.WebsocketEncodedMessage, error) {
	var (
		resp types.PriceResponse
	)

	h.logger.Debug("received message from Alpaca", zap.String("message", string(message)))

	// Check if the message is an authentication response
	var authResp []AuthenticatedMessage
	if err := json.Unmarshal(message, &authResp); err == nil && len(authResp) > 0 {
		for _, auth := range authResp {
			if auth.Type == "success" && auth.Msg == "authenticated" {
				h.logger.Info("received authentication response", zap.String("msg", auth.Msg))
				return resp, nil, nil
			}
		}
	}

	// Check if the message is a subscription confirmation
	var subscriptionConfirmation []SubscriptionConfirmationMessage
	if err := json.Unmarshal(message, &subscriptionConfirmation); err == nil && len(subscriptionConfirmation) > 0 {
		for _, subscription := range subscriptionConfirmation {
			if subscription.Type == "subscription" {
				h.logger.Info("subscription confirmed", zap.Strings("symbols", subscription.Quotes))
				return resp, nil, nil
			}
		}
	}

	// Check if the message is a quote update (array of quotes)
	var quotes []Quote
	if err := json.Unmarshal(message, &quotes); err == nil && len(quotes) > 0 {
		// This is a quote update
		h.logger.Debug("received quotes update", zap.Int("count", len(quotes)))
		resp, err := h.parseQuotes(quotes)
		if err != nil {
			return resp, nil, fmt.Errorf("failed to parse quotes: %w", err)
		}
		return resp, nil, nil
	}

	// If we got here, it's an unknown message type
	h.logger.Debug("received unknown message type", zap.String("message", string(message)))
	return resp, nil, nil
}

// CreateMessages is used to create messages to send to the data provider.
// For Alpaca, we first need to authenticate and then subscribe to quotes.
func (h *WebSocketHandler) CreateMessages(
	instruments []types.ProviderTicker,
) ([]handlers.WebsocketEncodedMessage, error) {
	// Store tickers in cache
	for _, ticker := range instruments {
		h.cache.Add(ticker)
	}

	// First message should be authentication
	authMsg, err := NewAuthMessage(h.apiKey, h.apiSecret)
	if err != nil {
		return nil, err
	}

	// Store tickers for subscription
	tickers := []string{}
	for _, ticker := range instruments {
		tickers = append(tickers, ticker.GetOffChainTicker())
	}

	// Create subscription messages
	subMsgs, err := h.NewSubscribeRequestMessage(tickers)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription messages: %w", err)
	}

	// Combine auth and subscription messages
	msgs := []handlers.WebsocketEncodedMessage{authMsg}
	msgs = append(msgs, subMsgs...)
	return msgs, nil
}

// HeartBeatMessages is used to send heartbeats to keep the connection alive.
// For Alpaca, we'll use a ping message if needed, and also handle re-auth and
// re-subscription if we haven't received messages in a while.
func (h *WebSocketHandler) HeartBeatMessages() ([]handlers.WebsocketEncodedMessage, error) {
	return nil, nil
}

// Copy is used to create a copy of the WebSocketHandler.
func (h *WebSocketHandler) Copy() types.PriceWebSocketDataHandler {
	return &WebSocketHandler{
		logger:    h.logger,
		ws:        h.ws,
		apiKey:    h.apiKey,
		apiSecret: h.apiSecret,
		cache:     types.NewProviderTickers(),
	}
}
