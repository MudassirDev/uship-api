package web

import (
	"encoding/json"
	"net/http"
)

func (apiCfg *apiConfig) handlerCheckout(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		PostalCode    string     `json:"postalCode"`
		Country       string     `json:"country"`
		ShippingPrice float64    `json:"shippingPrice"`
		Items         []CartItem `json:"items"`
	}

	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}
}
