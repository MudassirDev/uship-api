package web

import (
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

func (apiCfg *apiConfig) getRates(w http.ResponseWriter, r *http.Request) {
	_ = RouteAddress{
		Country:     apiCfg.PickupCountry,
		PostalCode:  apiCfg.PickupZip,
		AddressType: apiCfg.PickupType,
	}

	type Request struct {
		Country     string `json:"country"`
		PostalCode  string `json:"postalCode"`
		AddressType string `json:"type"`
	}
	var request Request

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	if ok := isValidAddress(request.AddressType); ok == false {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid address type %v", request.AddressType), "invalid address type")
		return
	}

	endpoint := fmt.Sprintf("%v/v2/raterequests?development=%v", apiCfg.BaseURL, ENV)

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
	}
	apiCfg.setHeaders(req)
	log.Println(req)
}

func isValidAddress(addressType string) bool {
	_, ok := validAddressTypes[addressType]
	return ok
}

