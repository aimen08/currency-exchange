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

		res, err := http.Get("https://v6.exchangerate-api.com/v6/1a007fbbc83055072c9e6479/latest/" + base)
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
		return c.JSON(rates)

	})

	app.Listen(":3000")
}
