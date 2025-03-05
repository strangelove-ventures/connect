# Alpaca WebSocket Provider

This provider implements a WebSocket connection with Alpaca's real-time market data API. Data is fed over the `wss://stream.data.alpaca.markets/v2/{feed}` URL. Feed can be either `iex`, `delayed_sip`, or `sip`. `delayed_sip` is on a 15 minute delay and is currently in use for our POC. `sip` requires a paid subscription while the others are free to use. 

## Authentication

The provider requires Alpaca API credentials (key and secret) to authenticate. When a connection is established, the provider must send an authentication message within **10 seconds**:

```json
{
  "action": "auth",
  "key": "YOUR_ALPACA_API_KEY",
  "secret": "YOUR_ALPACA_API_SECRET"
}
```

## Subscription

After authentication, the provider subscribes to quotes for the requested tickers within **10 seconds**:

```json
{ 
  "action": "subscribe",
  "quotes": ["AAPL", "MSFT", "GOOGL"]
}
```

Alpaca confirms the subscription with a response like this:

```json
[
  {
    "T": "subscription",
    "quotes": ["AAPL", "MSFT", "GOOGL"]
  }
]
```

## Data Format

Alpaca provides quotes in an array like so:

```json
[
  {
    "T":"q",
    "S":"AAPL",
    "bx":"Q",
    "bp":240.7,
    "bs":1,
    "ax":"V",
    "ap":240.76,
    "as":3,
    "c":["R"],
    "z":"C",
    "t":"2025-02-27T14:52:36.312631482Z"
  },
]
```

Where:
- `t`: Type (q for quote)
- `S`: Symbol
- `bx`: Bid exchange
- `bp`: Bid price
- `bs`: Bid size
- `ax`: Ask exchange
- `ap`: Ask price
- `as`: Ask size
- `c`: Conditions
- `z`: Tape
- `t`: Timestamp

## Usage

To use this provider, you need to:

1. Configure it with your Alpaca API credentials. Pass the credentials as environment variables `ALPACA_API_KEY` and `ALPACA_API_SECRET`.
2. Enable it in your configuration
3. Specify the tickers you want to monitor

The provider calculates the mid-price (average of bid and ask) for each quote received.

For more information, see [Alpaca's documentation](https://docs.alpaca.markets/docs/real-time-stock-pricing-data).
