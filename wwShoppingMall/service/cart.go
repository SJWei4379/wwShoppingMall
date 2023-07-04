package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"wwShoppingMall/globle"
	"wwShoppingMall/model"
	"wwShoppingMall/utils"
)

func AddCart(title, price, image string, currentID primitive.ObjectID) utils.Response {
	var cart model.Cart
	cart.Title = title
	cart.Price = price
	cart.Image = image
	cart.CurrentID = currentID
	if res, err := globle.CartColl.InsertOne(context.TODO(), cart); err != nil {
		return utils.ErrorMess("添加失败", nil)
	} else {
		return utils.SuccessMess("添加购物车成功", res)
	}

}

func GetCart() utils.Response {
	fs, err := globle.CartColl.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var data []bson.M
	if err = fs.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}
	return utils.SuccessMess("获取成功", data)
}

func DelCart(_id primitive.ObjectID) utils.Response {
	res, err := globle.CartColl.DeleteOne(context.TODO(), bson.M{"_id": _id})
	if err != nil {
		return utils.ErrorMess("删除失败", err.Error())
	}
	return utils.SuccessMess("删除成功", res)
}
