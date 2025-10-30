package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (apiCfg *apiConfig) handlerCheckout(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		PostalCode    string     `json:"postalCode"`
		Country       string     `json:"country"`
		ShippingPrice string     `json:"shippingPrice"`
		Items         []CartItem `json:"items"`
	}

	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}
	order := ShopifyOrder{
		DraftOrder: ShopifyDraftOrder{
			LineItems: requestBody.Items,
			ShippingAddress: ShopifyShippingAddress{
				Zip:     requestBody.PostalCode,
				Country: requestBody.Country,
			},
			ShippingLine: ShopifyShippingLine{
				Title: "UShip Shipping Price",
				Price: requestBody.ShippingPrice,
			},
		},
	}

	data, err := json.Marshal(order)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to create checkout")
		return
	}

	url := fmt.Sprintf("https://%v/admin/api/2025-10/draft_orders.json", apiCfg.ShopifyDomain)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to create checkout")
		return
	}

	req.Header.Add("X-Shopify-Access-Token", apiCfg.ShopifyAccessToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to create checkout")
		return
	}

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to create checkout")
		return
	}

	respondWithJSON(w, http.StatusCreated, string(resData))
}
