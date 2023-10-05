

# Exchange Rate API

This is a simple API that retrieves exchange rates from an external API and caches them using BigCache. The API is built using the Fiber web framework and retrieves exchange rates from the v6.exchangerate-api.com API.

## Installation

To install the API, clone the repository and run `go build` to build the binary. Then, run the binary to start the API server.

## Usage

The API exposes a single endpoint at `/exchange/:base`, where `:base` is the base currency code for the exchange rates to retrieve. For example, to retrieve exchange rates for USD, you would make a GET request to `/exchange/USD`.

The API caches exchange rates using BigCache, with a default cache time of 1440 minutes (24 hours). If exchange rates for a given base currency code are already in the cache, they will be retrieved from the cache instead of making a new request to the external API.

## Dependencies

This API uses the following dependencies:

- Fiber: A web framework for Go
- BigCache: An in-memory cache for Go

## Running with Air

You can run this code with `air`, a live-reloading tool for Go. To use `air`, simply install it using `go install github.com/cosmtrek/air@latest` and then run `air` in the same directory as your `main.go` file.

## API Result

The result of the API will be a JSON object containing the exchange rates for the specified base currency code. For example, if you make a GET request to `http://localhost:3000/exchange/USD`, the result might look like this:

```json
{
  "base_code": "USD",
  "conversion_rates": {
    "USD": 1,
    "EUR": 0.825,
    "CAD": 1.25
  }
}
```

Let me know if you have any questions or if there's anything else I can help you with!
