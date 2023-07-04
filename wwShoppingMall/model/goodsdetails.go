package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Goodsdetails struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	GoodsID  primitive.ObjectID `json:"goodsId" bson:"goodsId"`
	Title    string             `json:"title" bson:"title"`
	Details  string             `json:"details" bson:"details"`
	Price    string             `json:"price" bson:"price"`
	Topimage string             `json:"topimage" bson:"topimage"`
}
