package entity

type SwitchFilter struct {
	Type           string  `json:"type,omitempty" bson:"type,omitempty"`
	Manufacturer   string  `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`
	ActuationForce float64 `json:"actuation_force(g),omitempty" bson:"actuation_force(g),omitempty"`
}

type Switch struct {
	Name            string   `json:"name,omitempty"`
	Manufacturer    string   `json:"manufacturer,omitempty"`
	Type            string   `json:"type,omitempty"`
	StemType        string   `json:"stem_type,omitempty"`
	HousingColor    []string `json:"housing_color,omitempty"`
	StemColor       string   `json:"stem_color,omitempty"`
	RgbOption       bool     `json:"rgb_option,omitempty"`
	ActuationForce  float64  `json:"actuation_force(g),omitempty"`
	BottomOutForce  float64  `json:"bottom_out_force(g),omitempty"`
	ActuationLength float64  `json:"actuation_length(mm),omitempty"`
	TravelLength    float64  `json:"travel_length(mm),omitempty"`
	Picture         string   `json:"picture,omitempty"`
}
