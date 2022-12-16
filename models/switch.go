package model

type Switch struct {
	SwitchID        string         `json:"switch_id,omitempty" bson:"switch_id,omitempty"`
	Name            string         `json:"name,omitempty" bson:"name,omitempty"`
	Manufacturer    string         `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`
	Type            string         `json:"type,omitempty" bson:"type,omitempty"`
	StemType        string         `json:"stem_type,omitempty" bson:"stem_type,omitempty"`
	HousingColor    *HousingSwitch `json:"housing_color,omitempty" bson:"housing_color,omitempty"`
	StemColor       string         `json:"stem_color,omitempty" bson:"stem_color,omitempty"`
	RgbOption       bool           `json:"rgb_option,omitempty" bson:"rgb_option,omitempty"`
	ActuationForce  float64        `json:"actuation_force,omitempty" bson:"actuation_force,omitempty"`
	BottomOutForce  float64        `json:"bottom_out_force,omitempty" bson:"bottom_out_force,omitempty"`
	ActuationLength float64        `json:"actuation_length,omitempty" bson:"actuation_length,omitempty"`
	TravelLength    float64        `json:"travel_length,omitempty" bson:"travel_length,omitempty"`
	Picture         string         `json:"picture,omitempty" bson:"picture,omitempty"`
	CreatedAt       string         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt        string         `json:"updated_at,omitempty" bson:"update_at,omitempty"`
}

type HousingSwitch struct {
	TopHousing    string `json:"top_housing,omitempty" bson:"top_housing,omitempty"`
	BottomHousing string `json:"bottom_housing,omitempty" bson:"bottom_housing,omitempty"`
}
