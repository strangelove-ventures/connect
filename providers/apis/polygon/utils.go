package polygon

import (
	"time"

	"github.com/skip-mev/connect/v2/oracle/config"
)

const (
	// Name is the name of the Polygon.io provider.
	Name = "polygon_api"

	// URL is the base URL of the Polygon.io API.
	URL = "https://api.polygon.io"

	// Endpoint is the base URL of the Polygon.io API. This includes the base and quote
	// currency pairs that need to be inserted into the URL.
	//
	// Example from docs: https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/%s/%s
	// Last two '%s' are the date (Date must be formatted as "YYYY-MM-DD")
	Endpoint = "%s/v2/aggs/ticker/%s/range/1/day/%s/%s"
)

var DefaultAPIConfig = config.APIConfig{
	Name:             Name,
	Atomic:           false,
	Enabled:          true,
	Timeout:          5 * time.Second,
	Interval:         1 * time.Minute, // Adjust based on your API tier
	ReconnectTimeout: 5 * time.Second,
	MaxQueries:       1, // Polygon.io requires separate API calls per ticker
	Endpoints:        []config.Endpoint{{URL: URL}},
}

// PolygonResponse represents the response from Polygon.io API
type PolygonResponse struct {
	Ticker       string `json:"ticker"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []struct {
		Volume         float64 `json:"v"`
		VolumeWeighted float64 `json:"vw"`
		Open           float64 `json:"o"`
		Close          float64 `json:"c"`
		High           float64 `json:"h"`
		Low            float64 `json:"l"`
		Timestamp      int64   `json:"t"`
		NumberOfTrades int     `json:"n"`
	} `json:"results"`
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Count     int    `json:"count"`
}
