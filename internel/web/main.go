package web

import (
	"github.com/rs/cors"
	"net/http"
)

func CreateMux(apiKey string) http.Handler {
	mux := http.NewServeMux()
	apiCfg := apiConfig{
		APIKey:  apiKey,
		BaseURL: "https://api.ushipsandbox.com",
	}

	mux.HandleFunc("/getrates", apiCfg.getRates)

	return applyCors(mux)
}

func applyCors(mux *http.ServeMux) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://www.mystore.com"}, // Use your exact domain with scheme
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Authorization", "Content-Type"}, // Adjust as needed
		ExposedHeaders:   []string{"Content-Length"},                // Optional: expose only what's necessary
		AllowCredentials: true,                                      // Only allow if your frontend needs credentials (cookies, auth headers)
		MaxAge:           86400,                                     // Cache preflight response for 24 hours
	})

	handler := c.Handler(mux)
	return handler
}
