package alpaca

import (
	"encoding/json"
	"fmt"
	"math"

	connectmath "github.com/skip-mev/connect/v2/pkg/math"
	"github.com/skip-mev/connect/v2/providers/base/websocket/handlers"
)

type (
	// Action represents the action type for Alpaca messages
	Action string
)

const (
	// AuthAction is the action for authentication
	AuthAction Action = "auth"

	// SubscribeAction is the action for subscribing to market data
	SubscribeAction Action = "subscribe"

	// QuotesStream is used to subscribe to quotes
	QuotesStream = "quotes"
)

// AuthMessage represents the authentication message to be sent to Alpaca
type AuthMessage struct {
	Action Action `json:"action"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type AuthenticatedMessage struct {
	Type string `json:"T"`
	Msg  string `json:"msg"`
}

// SubscribeMessage represents a message to subscribe to market data
type SubscribeMessage struct {
	Action Action   `json:"action"`
	Quotes []string `json:"quotes,omitempty"`
}

// SubscriptionConfirmationMessage represents a subscription confirmation message from Alpaca
type SubscriptionConfirmationMessage struct {
	Type   string   `json:"T"`
	Quotes []string `json:"quotes,omitempty"`
}

// Quote represents a quote response from Alpaca
type Quote struct {
	Type        string   `json:"T"`
	Symbol      string   `json:"S"`
	BidExchange string   `json:"bx"`
	BidPrice    float64  `json:"bp"`
	BidSize     float64  `json:"bs"`
	AskExchange string   `json:"ax"`
	AskPrice    float64  `json:"ap"`
	AskSize     float64  `json:"as"`
	Conditions  []string `json:"c"`
	Tape        string   `json:"z"`
	Timestamp   string   `json:"t"`
}

// NewAuthMessage creates a new authentication message with the provided API key and secret
func NewAuthMessage(apiKey, apiSecret string) ([]byte, error) {
	msg := AuthMessage{
		Action: AuthAction,
		Key:    apiKey,
		Secret: apiSecret,
	}

	bz, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal auth message: %w", err)
	}

	return bz, nil
}

// NewSubscribeRequestMessage returns a new subscribe request message for the given instruments
func (h *WebSocketHandler) NewSubscribeRequestMessage(instruments []string) ([]handlers.WebsocketEncodedMessage, error) {
	numInstruments := len(instruments)
	if numInstruments == 0 {
		return nil, fmt.Errorf("no instruments provided")
	}

	numBatches := int(math.Ceil(float64(numInstruments) / float64(h.ws.MaxSubscriptionsPerBatch)))
	msgs := make([]handlers.WebsocketEncodedMessage, numBatches)
	for i := 0; i < numBatches; i++ {
		// Get the instruments for the batch.
		start := i * h.ws.MaxSubscriptionsPerBatch
		end := connectmath.Min((i+1)*h.ws.MaxSubscriptionsPerBatch, numInstruments)

		bz, err := json.Marshal(SubscribeMessage{
			Action: SubscribeAction,
			Quotes: instruments[start:end],
		})
		if err != nil {
			return nil, fmt.Errorf("failed to marshal subscribe message: %w", err)
		}
		msgs[i] = bz
	}

	return msgs, nil
}
