package web

type apiConfig struct {
	APIKey        string
	BaseURL       string
	PickupCountry string
	PickupZip     string
	PickupType    string
}

type RateRequestBody struct {
	Route Route     `json:"route"`
	Items []LtlItem `json:"items"`
}

type Route struct {
	RouteItems []RouteItem `json:"items"`
}

type RouteItem struct {
	Accessorials []string     `json:"accessorials,omitempty"`
	Address      RouteAddress `json:"address"`
}

type RouteAddress struct {
	Country     string `json:"country"`
	PostalCode  string `json:"postalCode"`
	AddressType string `json:"type"`
}

type LtlItem struct {
	Commodity      string  `json:"commodity"`
	Description    string  `json:"description"`
	HandlingUnit   string  `json:"handlingUnit"`
	Title          string  `json:"title"`
	Packaging      string  `json:"packaging"`
	WeightInGrams  float64 `json:"weightInGrams"`
	WidthInMeters  float64 `json:"widthInMeters"`
	HeightInMeters float64 `json:"heightInMeters"`
	LengthInMeters float64 `json:"lengthInMeters"`
	Hazardous      bool    `json:"hazardous"`
	Stackable      bool    `json:"stackable"`
	UnitCount      int     `json:"unitCount"`
}
