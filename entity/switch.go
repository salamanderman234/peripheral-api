package entity

import (
	model "github.com/salamanderman234/peripheral-api/models"
)

type Switch struct {
	SwitchID        string               `json:"switch_id,omitempty" query:"switch_id"`
	Name            string               `json:"name,omitempty" query:"name"`
	Manufacturer    string               `json:"manufacturer,omitempty" query:"manufacturer"`
	Type            string               `json:"type,omitempty" query:"type"`
	StemType        string               `json:"stem_type,omitempty" query:"stem_type"`
	HousingColor    *model.HousingSwitch `json:"housing_color,omitempty"`
	StemColor       string               `json:"stem_color,omitempty" query:"stem_color"`
	RgbOption       bool                 `json:"rgb_option,omitempty" query:"rgb_option"`
	ActuationForce  float64              `json:"actuation_force,omitempty" query:"actuation_force"`
	BottomOutForce  float64              `json:"bottom_out_force,omitempty" query:"bottom_force"`
	ActuationLength float64              `json:"actuation_length,omitempty" query:"actuation_length"`
	TravelLength    float64              `json:"travel_length,omitempty" query:"travel_length"`
	Picture         string               `json:"picture,omitempty"`
}

type SwitchPolicy struct {
	Data                *Switch `json:"data"`
	NameMessage         string  `json:"name,omitempty"`
	ManufacturerMessage string  `json:"manufacturer,omitempty"`
	TypeMessage         string  `json:"type,omitempty"`
	StemTypeMessage     string  `json:"stem_type,omitempty"`
	StemColorMessage    string  `json:"stem_color,omitempty"`
}

type SwitchUpdate struct {
	Filter      Switch `json:"filter"`
	UpdateField Switch `json:"update_field"`
}

func (s *Switch) SwitchInsertValidation(similarSwitch []Switch) *SwitchPolicy {
	var policy SwitchPolicy
	isReturn := false

	policy.Data = s

	if s.Name == "" {
		policy.NameMessage = "Name is required"
		isReturn = true
	} else if len(similarSwitch) != 0 {
		policy.NameMessage = "Name must be unique"
		isReturn = true
	}

	if s.Manufacturer == "" {
		policy.ManufacturerMessage = "Manufacturer is required"
		isReturn = true
	}

	if s.Type == "" {
		policy.TypeMessage = "Type is required"
		isReturn = true
	} else {
		if s.Type != "linear" && s.Type != "clicky" && s.Type != "silent" {
			policy.TypeMessage = "Type must be either linear, clicky or silent"
			isReturn = true
		}
	}

	if s.StemType == "" {
		policy.StemTypeMessage = "Stem type is required"
		isReturn = true
	}

	if s.StemColor == "" {
		policy.StemColorMessage = "Stem color is required"
		isReturn = true
	}

	if s.HousingColor.TopHousing == "" || s.HousingColor.BottomHousing == "" {
		policy.StemColorMessage = "Housing color top housing and bottom housing is required"
		isReturn = true
	}

	if isReturn {
		return &policy
	}
	return nil
}

func (s *Switch) SwitchUpdateValidation(similarSwitch []Switch) *SwitchPolicy {
	var policy SwitchPolicy
	isReturn := false
	policy.Data = s
	if len(similarSwitch) != 0 && s.Name != "" {
		policy.NameMessage = "Name must be unique"
		isReturn = true
	}
	if (s.Type != "") && (s.Type != "linear" && s.Type != "clicky" && s.Type != "silent") {
		policy.TypeMessage = "Type must be either linear, clicky or silent"
		isReturn = true
	}

	if isReturn {
		return &policy
	}
	return nil
}
