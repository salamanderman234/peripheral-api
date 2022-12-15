package entity

type Switch struct {
	Name            string   `json:"name,omitempty"`
	Manufacturer    string   `json:"manufacturer,omitempty" query:"manufacturer"`
	Type            string   `json:"type,omitempty" query:"type"`
	StemType        string   `json:"stem_type,omitempty"`
	HousingColor    []string `json:"housing_color,omitempty"`
	StemColor       string   `json:"stem_color,omitempty"`
	RgbOption       bool     `json:"rgb_option,omitempty"`
	ActuationForce  float64  `json:"actuation_force(g),omitempty" query:"actuation_force"`
	BottomOutForce  float64  `json:"bottom_out_force(g),omitempty"`
	ActuationLength float64  `json:"actuation_length(mm),omitempty"`
	TravelLength    float64  `json:"travel_length(mm),omitempty"`
	Picture         string   `json:"picture,omitempty"`
}
