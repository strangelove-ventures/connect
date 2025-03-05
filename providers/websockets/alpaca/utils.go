package alpaca

import (
	"github.com/skip-mev/connect/v2/oracle/config"
)

const (
	// Name is the name of the Alpaca provider.
	Name = "alpaca_ws"

	// URL is the production Alpaca Websocket URL.
	// This is for the delayed SIP feed. Switch from
	// `delayed_sip` to `sip` for real-time. This requires
	// a paid API token.
	WSS = "wss://stream.data.alpaca.markets/v2/delayed_sip"
)

// DefaultWebSocketConfig is the default configuration for the Alpaca Websocket.
var DefaultWebSocketConfig = config.WebSocketConfig{
	Name:                          Name,
	Enabled:                       true,
	MaxBufferSize:                 config.DefaultMaxBufferSize,
	ReconnectionTimeout:           config.DefaultReconnectionTimeout,
	PostConnectionTimeout:         config.DefaultPostConnectionTimeout,
	HandshakeTimeout:              config.DefaultHandshakeTimeout,
	Endpoints:                     []config.Endpoint{{URL: WSS}},
	ReadBufferSize:                config.DefaultReadBufferSize,
	WriteBufferSize:               config.DefaultWriteBufferSize,
	EnableCompression:             config.DefaultEnableCompression,
	ReadTimeout:                   config.DefaultReadTimeout,
	WriteTimeout:                  config.DefaultWriteTimeout,
	PingInterval:                  config.DefaultPingInterval,
	WriteInterval:                 config.DefaultWriteInterval,
	MaxReadErrorCount:             config.DefaultMaxReadErrorCount,
	MaxSubscriptionsPerConnection: config.DefaultMaxSubscriptionsPerConnection,
	MaxSubscriptionsPerBatch:      config.DefaultMaxSubscriptionsPerBatch,
}
