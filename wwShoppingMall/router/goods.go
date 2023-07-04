package router

import (
	"github.com/gin-gonic/gin"
	"wwShoppingMall/controller"
)

func GoodsRouter(engine *gin.Engine) {
	goods := engine.Group("goods")
	{
		goods.GET("/get", controller.GetGoods)
		goods.GET("/find", controller.FindGoods)
		goods.GET("/details", controller.FindGoodsDetails)
		goods.GET("/buyfind", controller.BuyFindGoods)
	}
}
