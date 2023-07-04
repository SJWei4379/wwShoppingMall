package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Goods struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
	Price string             `json:"price" bson:"price"`
	Url   string             `json:"url" bson:"url"`
	Tag   string             `json:"tag" bson:"tag"`
}
