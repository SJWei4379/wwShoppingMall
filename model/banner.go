package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Banner struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Url   string             `json:"url" bson:"url"`
	Descs string             `json:"descs" bson:"descs"`
}
