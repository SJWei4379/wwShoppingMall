package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"wwShoppingMall/globle"
	"wwShoppingMall/utils"
)

func GetKeywords() utils.Response {
	fs, err := globle.Keywords.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var data []bson.M
	if err = fs.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}
	return utils.SuccessMess("获取成功", data)
}
