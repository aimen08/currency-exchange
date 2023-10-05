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

## License

This code is licensed under the MIT license. See the LICENSE file for more information.

Let me know if you have any questions or if there's anything else I can help you with!

aimen08: add "you can run it with air " and also add the result of the api

GitHub Copilot: Sure, here's the updated code with the requested changes:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
)

type SelectedRates struct {
	USD float64 `json:"USD"`
	EUR float64 `json:"EUR"`
	CAD float64 `json:"CAD"`
}

type Rates struct {
	Base  string        `json:"base_code"`
	Rates SelectedRates `json:"conversion_rates"`
}

func main() {
	app := fiber.New()

	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(1440*time.Minute))

	app.Get("/exchange/:base", func(c *fiber.Ctx) error {
		base := c.Params("base")

		entry, _ := cache.Get(base)

		if len(entry) > 0 {
			cached := Rates{}
			parseErrCache := json.Unmarshal(entry, &cached)
			if parseErrCache != nil {
				return parseErrCache
			}
			fmt.Println("From cache")
			return c.JSON(cached)
		}

		res, err := http.Get("https://v6.exchangerate-api.com/v6/YOUR-API-KEY/latest/" + base)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		rates := Rates{}
		parseErr := json.Unmarshal(body, &rates)
		if parseErr != nil {
			return parseErr
		}

		cache.Set(base, body)

		fmt.Println("From API")
		return c.JSON(rates)
	})

	app.Listen(":3000")
}
```

You can run this code with `air`, a live-reloading tool for Go. To use `air`, simply install it using `go install github.com/cosmtrek/air@latest` and then run `air` in the same directory as your `main.go` file.

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

aimen08: write readme for the code in the main.go file

add "you can run it with air " and also add the result of the api



GitHub Copilot: Sure, here's a sample README file for the code in the `main.go` file:

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
