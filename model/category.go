package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tag   string             `json:"tag" bson:"tag"`
	Image string             `json:"image" bson:"image"`
	Cate  string             `json:"cate" bson:"cate"`
}
