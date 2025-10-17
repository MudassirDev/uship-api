package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var validAddressTypes = map[string]struct{}{
	"Residence":                            {},
	"BusinessWithLoadingDockOrForklift":    {},
	"BusinessWithoutLoadingDockOrForklift": {},
	"ConstructionSite":                     {},
	"TradeShowOrConvention":                {},
	"Port":                                 {},
	"StorageFacility":                      {},
	"MilitaryBase":                         {},
	"Airport":                              {},
	"OtherSecuredLocation":                 {},
	"GovernmentLocation":                   {},
	"FarmRanchEstate":                      {},
	"ReligiousInstitution":                 {},
	"GolfCourseResortPark":                 {},
	"DistributionCenter":                   {},
	"Auction":                              {},
}

type RateRequestStruct struct {
	RouteAddress
	Items []LtlItem `json:"items"`
}

func (apiCfg *apiConfig) getRates(w http.ResponseWriter, r *http.Request) {
	pickupAddress := RouteAddress{
		Country:     apiCfg.PickupCountry,
		PostalCode:  apiCfg.PickupZip,
		AddressType: apiCfg.PickupType,
	}

	var request RateRequestStruct

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	if ok := validateRequestStruct(&request); !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"), "invalid payload")
		return
	}
	endpoint := fmt.Sprintf("%v/v2/raterequests?development=%v", apiCfg.BaseURL, ENV)

	payload, err := createPayload(&request, &pickupAddress)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to create rates")
		return
	}

	req, err := http.NewRequest("POST", endpoint, payload)
	if err != nil {
	}
	apiCfg.setHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to create rates")
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "error while decoding response")
		return
	}

	if res.StatusCode != 201 {
		respondWithError(w, http.StatusInternalServerError, fmt.Errorf("reponse: %v, res %v", string(data), res), "failed to create rates")
		return
	}

	rates, err := apiCfg.fetchRates(res.Header.Get("location"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err, "failed to fetch rates")
		return
	}

	respondWithJSON(w, http.StatusOK, rates)
}

func (apiCfg *apiConfig) fetchRates(url string) (any, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	apiCfg.setHeaders(request)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	var response any
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()

	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func createPayload(requestStruct *RateRequestStruct, pickupAddress *RouteAddress) (*bytes.Buffer, error) {
	data := RateRequestBody{
		Items: requestStruct.Items,
	}
	futureDate := time.Now().AddDate(0, 1, 0).Format("2006/01/02")

	route := Route{
		RouteItems: []RouteItem{
			{
				Address: *pickupAddress,
				TimeFrameValue: &TimeFrame{
					EarliestArrival: futureDate,
					LatestArrival: futureDate,
					TimeFrameType: "on",
				},
			},
			{
				Address:      requestStruct.RouteAddress,
				Accessorials: []string{"LiftgateRequired"},
			},
		},
	}
	data.Route = route

	payload, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(payload), nil
}

func validateRequestStruct(requestStruct *RateRequestStruct) bool {
	if requestStruct.Country == "" {
		return false
	}
	if requestStruct.AddressType == "" {
		return false
	}
	if requestStruct.PostalCode == "" {
		return false
	}
	if len(requestStruct.Items) == 0 {
		return false
	}
	return isValidAddress(requestStruct.AddressType)
}

func isValidAddress(addressType string) bool {
	_, ok := validAddressTypes[addressType]
	return ok
}
