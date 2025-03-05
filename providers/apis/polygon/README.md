# Polygon.io API Provider

This provider implements an HTTP API connection with [Polygon.io's](https://polygon.io/) market data API. Data is fetched from the REST API endpoint `https://api.polygon.io/v2/aggs/ticker/{ticker}/range/1/day/{fromDate}/{toDate}`.

## Authentication

The provider requires a Polygon.io API key to authenticate. The API key is passed as a Bearer token in the Authorization header:

```
Authorization: Bearer YOUR_POLYGON_API_KEY
```

To use this provider, set the `POLYGON_API_KEY` environment variable with your Polygon.io API key.

## Data Format

Polygon.io provides aggregated bar data in the following format:

```json
{
  "ticker": "AAPL",
  "queryCount": 1,
  "resultsCount": 1,
  "adjusted": true,
  "results": [
    {
      "v": 76286780,        // Volume
      "vw": 174.3024,       // Volume weighted average price
      "o": 173.44,          // Open price
      "c": 175.05,          // Close price
      "h": 175.1,           // High price
      "l": 173.17,          // Low price
      "t": 1677887999999,   // Timestamp (Unix ms)
      "n": 538470           // Number of transactions
    }
  ],
  "status": "OK",
  "request_id": "6a7e466b6f1c51c96c4751c15450fa59",
  "count": 1
}
```

The provider uses the average of the previous day's open, close, high, and low price.

## Usage

To use this provider, you need to:

1. Sign up for a Polygon.io account and obtain an API key
2. Set the `POLYGON_API_KEY` environment variable
3. Enable the provider in your configuration
4. Specify the tickers you want to monitor

For more information, see [Polygon.io's documentation](https://polygon.io/docs/stocks/getting-started).
