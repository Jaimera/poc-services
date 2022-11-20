package dto

// PortDto is a data transfer object, to be used on API and integrations
type PortDto struct {
	ID        uint32   `json:"id"`
	Slug      string   `json:"slug"`
	Name      string   `json:"name"`
	City      string   `json:"city"`
	Province  string   `json:"province"`
	Country   string   `json:"country"`
	Alias     []string `json:"alias"`
	Regions   []string `json:"regions"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Timezone  string   `json:"timezone"`
	Unlocs    []string `json:"unlocs"`
	Code      *string  `json:"code"`
}
