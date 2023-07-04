package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"wwShoppingMall/globle"
	"wwShoppingMall/utils"
)

func FindGoodsDetails(id primitive.ObjectID) utils.Response {
	var res bson.M
	err := globle.Goodsdetails.FindOne(context.TODO(), bson.M{"goodsId": id}).Decode(&res)
	if err != nil {
		var data bson.M
		err := globle.Goodsdetails.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
		if err != nil {
			return utils.ErrorMess("查询失败", nil)
		}
		return utils.SuccessMess("成功", data)
	}
	return utils.SuccessMess("成功", res)

	//if res != nil {
	//	return utils.SuccessMess("成功", res)
	//} else {
	//	var data bson.M
	//	err := globle.Goodsdetails.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
	//	if err != nil {
	//		return utils.ErrorMess("查询失败", nil)
	//	}
	//	return utils.SuccessMess("成功", data)
	//}

}
