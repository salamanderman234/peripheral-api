package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Switch struct {
	ID              primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty"`
	Manufacturer    string             `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`
	Type            string             `json:"type,omitempty" bson:"type,omitempty"`
	StemType        string             `json:"stem_type,omitempty" bson:"stem_type,omitempty"`
	HousingColor    []string           `json:"housing_color,omitempty" bson:"housing_color,omitempty"`
	StemColor       string             `json:"stem_color,omitempty" bson:"stem_color,omitempty"`
	RgbOption       bool               `json:"rgb_option,omitempty" bson:"rgb_option,omitempty"`
	ActuationForce  float64            `json:"actuation_force(g),omitempty" bson:"actuation_force(g),omitempty"`
	BottomOutForce  float64            `json:"bottom_out_force(g),omitempty" bson:"bottom_out_force(g),omitempty"`
	ActuationLength float64            `json:"actuation_length(mm),omitempty" bson:"actuation_length(mm),omitempty"`
	TravelLength    float64            `json:"travel_length(mm),omitempty" bson:"travel_length(mm),omitempty"`
	Picture         string             `json:"picture,omitempty" bson:"picture,omitempty"`
	CreatedAt       time.Time          `json:"-" bson:"created_at,omitempty"`
}
