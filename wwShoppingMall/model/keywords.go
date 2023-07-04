package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Keywords struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string             `json:"content" bson:"content"`
}
