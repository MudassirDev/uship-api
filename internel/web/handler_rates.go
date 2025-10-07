package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	log.Println(res)
}

func createPayload(requestStruct *RateRequestStruct, pickupAddress *RouteAddress) (*bytes.Buffer, error) {
	data := RateRequestBody{
		Items: requestStruct.Items,
	}

	route := Route{
		RouteItems: []RouteItem{
			{
				Address: *pickupAddress,
			},
			{
				Address:      requestStruct.RouteAddress,
				Accessorials: []string{"Inside", "LiftgateRequired"},
			},
		},
	}
	data.Route = route

	payload, err := json.Marshal(data)
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
