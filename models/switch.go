package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Switch struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Type      string             `bson:"type,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}
