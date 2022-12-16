package entity

import model "github.com/salamanderman234/peripheral-api/models"

type Switch struct {
	SwitchID     string `json:"switch_id,omitempty" query:"switch_id"`
	Name         string `json:"name,omitempty" query:"name"`
	Manufacturer string `json:"manufacturer,omitempty" query:"manufacturer"`
	Type         string `json:"type,omitempty" query:"type"`
	StemType     string `json:"stem_type,omitempty"`
	// harus diginikan agar bisa kosong secara default saat json
	HousingColor    *model.HousingSwitch `json:"housing_color,omitempty"`
	StemColor       string               `json:"stem_color,omitempty"`
	RgbOption       bool                 `json:"rgb_option,omitempty"`
	ActuationForce  float64              `json:"actuation_force,omitempty" query:"actuation_force"`
	BottomOutForce  float64              `json:"bottom_out_force,omitempty"`
	ActuationLength float64              `json:"actuation_length,omitempty"`
	TravelLength    float64              `json:"travel_length,omitempty"`
	Picture         string               `json:"picture,omitempty"`
}

type SwitchUpdate struct {
	Filter      Switch `json:"filter"`
	UpdateField Switch `json:"update_field"`
}
