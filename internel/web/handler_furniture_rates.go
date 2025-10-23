package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) getFurnitureRates(w http.ResponseWriter, r *http.Request) {
	endpoint := "/v2/fixedprice"
	url := apiCfg.BaseURL + endpoint

	type Request struct {
		OriginPostalCode string          `json:"originPostalCode"`
		Items            []FurnitureItem `json:"items"`
	}

	var request Request
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, err, "wrong payload")
		return
	}

	var requestBody = FurnitureRequestBody{
		OriginPostalCode:      request.OriginPostalCode,
		DestinationPostalCode: apiCfg.PickupZip,
		Source:                "API",
		Items:                 request.Items,
	}

	data, err := json.MarshalIndent(requestBody, "", "")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "error while making request")
		return
	}
	payload := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "error while making request")
		return
	}
	apiCfg.setHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to create rates")
		return
	}

	if res.StatusCode != 200 {
		respondWithError(w, http.StatusInternalServerError, fmt.Errorf("reason: %v, code: %v", res.Status, res.StatusCode), "failed to create rates")
		return
	}

	var response any
	newDecoder := json.NewDecoder(res.Body)
	defer res.Body.Close()

	if err := newDecoder.Decode(&response); err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to decode response")
		return
	}

	respondWithJSON(w, http.StatusOK, response)
}
