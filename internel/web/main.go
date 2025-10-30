package web

import (
	"net/http"

	"github.com/rs/cors"
)

var (
	ENV           string
	isDevelopment bool
)

func CreateMux(apiKey, host, env, pickupCountry, pickupZip, pickupType, shopifyDomain, shopifyAccessToken string) http.Handler {
	ENV = env
	isDevelopment = env == "development"

	mux := http.NewServeMux()
	apiCfg := apiConfig{
		APIKey:        apiKey,
		BaseURL:       "https://api.ushipsandbox.com",
		PickupCountry: pickupCountry,
		PickupZip:     pickupZip,
		PickupType:    pickupType,
		ShopifyDomain: shopifyDomain,
		ShopifyAccessToken: shopifyAccessToken,
	}
	if !isDevelopment {
		apiCfg.BaseURL = "https://api.uship.com"
	}

	mux.HandleFunc("POST /get-rates", apiCfg.getRates)
	mux.HandleFunc("POST /get-furniture-rates", apiCfg.getFurnitureRates)
	mux.HandleFunc("POST /get-checkout-url", apiCfg.handlerCheckout)

	return applyCors(mux, host)
}

func applyCors(mux *http.ServeMux, host string) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{host},
		MaxAge:         86400,
	})

	handler := c.Handler(mux)
	return handler
}
