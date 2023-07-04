package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"wwShoppingMall/globle"
	"wwShoppingMall/utils"
)

func Category(tag string) utils.Response {
	fs, err := globle.CategoryColl.Find(context.TODO(), bson.M{"cate": tag})
	if err != nil {
		log.Fatal(err)
	}
	var data []bson.M
	if err = fs.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}
	if data == nil {
		return utils.ErrorMess("获取失败", nil)
	}
	return utils.SuccessMess("获取成功", data)

	//var data bson.M
	//err := globle.CategoryColl.FindOne(context.TODO(), bson.M{"cate": cate}).Decode(&data)
	//if err != nil {
	//	return utils.ErrorMess("查询失败", nil)
	//}
	//return utils.SuccessMess("成功", data)
}
