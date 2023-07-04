package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"wwShoppingMall/globle"
	"wwShoppingMall/utils"
)

func GetGoods(currPage string) utils.Response {
	//size,err := strconv.ParseInt(pageSize,10,64)
	//if err != nil {
	//	return utils.ErrorMess("行数字段过长",size)
	//}
	curr, err := strconv.ParseInt(currPage, 10, 64)
	if err != nil {
		return utils.ErrorMess("指定页面字段过长", nil)
	}
	var size int64
	size = 10
	skip := (curr - 1) * size
	opt := options.FindOptions{
		Limit: &size,
		Skip:  &skip,
		Sort:  bson.M{"_id": 1},
	}
	cur, err := globle.GoodsColl.Find(context.TODO(), bson.M{}, &opt)
	if err != nil {
		return utils.ErrorMess("查询失败", err.Error())
	}
	var data []bson.M
	if err := cur.All(context.TODO(), &data); err != nil {
		return utils.ErrorMess("没有数据", err.Error())
	}
	if data == nil {
		return utils.ErrorMess("数据为空", nil)
	}
	return utils.SuccessMess("查询成功", data)
}

func FindGoods(search string) utils.Response {

	regex := bson.M{"$regex": primitive.Regex{Pattern: search, Options: "i"}}
	filter := bson.M{
		"$or": bson.A{
			bson.M{"title": regex},
			bson.M{"tag": regex},
		},
	}

	//filter := bson.M{
	//	"$or": bson.M{
	//		bson.M{"title": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "i"}}},
	//		bson.M{"tag": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "i"}}},
	//	},
	//}

	fs, err := globle.GoodsColl.Find(context.TODO(), filter)
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

	//var goods model.Goods
	//err := globle.GoodsColl.FindOne(context.TODO(), filter).Decode(&goods)
	//if err != nil {
	//	return utils.ErrorMess("查询失败", nil)
	//}
	//fmt.Printf("found document %v", goods)
	//return utils.SuccessMess("成功", goods)
}

func BuyFindGoods(_id primitive.ObjectID) utils.Response {
	var res bson.M
	err := globle.Goodsdetails.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&res)
	if err != nil {
		return utils.ErrorMess("查询失败", nil)
	}
	return utils.SuccessMess("成功", res)
}
