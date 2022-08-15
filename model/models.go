package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BSON model for MongoDB service
type Message struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User string             `json:"user,omitempty"`
	Text string             `json:"text,omitempty"`
	Date string             `json:"date,omitempty"`
}
