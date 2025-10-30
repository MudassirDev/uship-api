package web

type apiConfig struct {
	APIKey             string
	BaseURL            string
	PickupCountry      string
	PickupZip          string
	PickupType         string
	ShopifyDomain      string
	ShopifyAccessToken string
}

type RateRequestBody struct {
	Route Route     `json:"route"`
	Items []LtlItem `json:"items"`
}

type Route struct {
	RouteItems []RouteItem `json:"items"`
}

type RouteItem struct {
	Accessorials   []string     `json:"accessorials,omitempty"`
	Address        RouteAddress `json:"address"`
	TimeFrameValue *TimeFrame   `json:"timeFrame,omitempty"`
}

type TimeFrame struct {
	EarliestArrival string `json:"earliestArrival"`
	LatestArrival   string `json:"latestArrival"`
	TimeFrameType   string `json:"timeFrameType"`
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

type FurnitureRequestBody struct {
	DestinationPostalCode string          `json:"destinationPostalCode"`
	OriginPostalCode      string          `json:"originPostalCode"`
	Source                string          `json:"source"`
	Items                 []FurnitureItem `json:"items"`
}

type FurnitureItem struct {
	HeightInMeters float64 `json:"heightInMeters"`
	LengthInMeters float64 `json:"lengthInMeters"`
	WeightInGrams  float64 `json:"weightInGrams"`
	WidthInMeters  float64 `json:"widthInMeters"`
	UnitCount      int     `json:"unitCount"`
}
type ShopifyOrder struct {
	DraftOrder ShopifyDraftOrder `json:"draft_order"`
}
type ShopifyDraftOrder struct {
	LineItems       []CartItem             `json:"line_items"`
	ShippingLine    ShopifyShippingLine    `json:"shipping_line"`
	ShippingAddress ShopifyShippingAddress `json:"shipping_address"`
}
type ShopifyShippingAddress struct {
	Country string `json:"country"`
	Zip     string `json:"zip"`
}
type ShopifyShippingLine struct {
	Title string `json:"title"`
	Price string `json:"price"`
}
type CartItem struct {
	Quantity   int                `json:"quantity"`
	VariantID  int                `json:"variant_id"`
	Properties []CartItemProperty `json:"properties"`
}

type CartItemProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
