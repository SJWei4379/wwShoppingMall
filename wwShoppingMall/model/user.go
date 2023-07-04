package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//type User struct {
//	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
//	Name      string             `json:"name" bson:"name"`
//	Account   string             `json:"account" bson:"account"`
//	Password  string             `json:"password" bson:"password"`
//	Age       int64              `json:"age" bson:"age"`
//	Salt      string             `json:"salt,omitempty" bson:"salt,omitempty"`
//	AvatarUrl string             `bson:"avatarUrl" json:"avatarUrl"` //头像地址
//	Sex       string             `bson:"sex" json:"sex"`
//	Phone     string             `bson:"phone" json:"phone"`
//	OpenId    string             `bson:"openId" json:"openId"`
//}

type UserWX struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OpenID     string             `json:"openId" bson:"openId"`
	SessionKey string             `json:"session_key" bson:"session_key"`
}
