package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Image     string             `json:"image" bson:"image"`
	Price     string             `json:"price" bson:"price"`
	CurrentID primitive.ObjectID `json:"currentID" bson:"currentID"`
}
