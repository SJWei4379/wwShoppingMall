package globle

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient  *mongo.Client
	UserColl     *mongo.Collection
	UserWXColl   *mongo.Collection
	LogColl      *mongo.Collection
	BannerColl   *mongo.Collection
	CartColl     *mongo.Collection
	CategoryColl *mongo.Collection
	GoodsColl    *mongo.Collection
	Goodsdetails *mongo.Collection
	Keywords     *mongo.Collection
)

//MySql
//var (
//	DB *sqlx.DB
//)
