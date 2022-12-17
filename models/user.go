package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	CreatedAt string             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt  string             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
