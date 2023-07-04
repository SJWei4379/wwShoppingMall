package router

import (
	"github.com/gin-gonic/gin"
	"wwShoppingMall/controller"
)

func CartRouter(engine *gin.Engine) {
	cart := engine.Group("cart")
	{
		cart.GET("/add", controller.AddCart)
		cart.GET("/get", controller.GetCart)
		cart.GET("/del", controller.DelCart)
	}
}
