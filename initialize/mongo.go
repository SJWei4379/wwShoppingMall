package initialize

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"wwShoppingMall/globle"
)

func MongoInit() {
	if globle.MongoClient == nil {
		globle.MongoClient = getMongoClient("mongodb://localhost:27017")
	}
	ShoppingMall := globle.MongoClient.Database("wwShoppingMall")
	{
		globle.UserColl = ShoppingMall.Collection("user")
		globle.UserWXColl = ShoppingMall.Collection("userwx")
		globle.LogColl = ShoppingMall.Collection("log")
		globle.BannerColl = ShoppingMall.Collection("banner")
		globle.CartColl = ShoppingMall.Collection("cart")
		globle.CategoryColl = ShoppingMall.Collection("category")
		globle.GoodsColl = ShoppingMall.Collection("goods")
		globle.Goodsdetails = ShoppingMall.Collection("goodsdetails")
		globle.Keywords = ShoppingMall.Collection("keywords")
	}
}

func getMongoClient(url string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(url)

	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	if err = MongoClient.Ping(context.TODO(), nil); err != nil {
		fmt.Println(err)
	}
	return MongoClient
}

// Mysql
//func MysqlInit() {
//	if globle.DB == nil {
//		globle.DB = getMysqlClient("root:root@tcp(localhost:3306)/goTest")
//	}
//
//}
//
//func getMysqlClient(url string) *sqlx.DB {
//	DB, _ := sqlx.Open("mysql", url)
//	//设置数据库最大连接数
//	DB.SetConnMaxLifetime(100)
//	//设置上数据库最大闲置连接数
//	DB.SetMaxIdleConns(10)
//	//验证连接
//	if err := DB.Ping(); err != nil {
//		fmt.Println("open database fail")
//		return nil
//	}
//	fmt.Println("connection success")
//	return DB
//}
