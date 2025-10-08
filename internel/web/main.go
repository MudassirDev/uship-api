package web

import (
	"net/http"

	"github.com/rs/cors"
)

var (
	ENV           string
	isDevelopment bool
)

func CreateMux(apiKey, host, env, pickupCountry, pickupZip, pickupType string) http.Handler {
	ENV = env
	isDevelopment = env == "development"

	mux := http.NewServeMux()
	apiCfg := apiConfig{
		APIKey:        apiKey,
		BaseURL:       "https://api.ushipsandbox.com",
		PickupCountry: pickupCountry,
		PickupZip:     pickupZip,
		PickupType:    pickupType,
	}
	if !isDevelopment {
		apiCfg.BaseURL = "https://api.uship.com"
	}

	mux.HandleFunc("POST /getrates", apiCfg.getRates)

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
