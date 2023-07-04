package router

import (
	"github.com/gin-gonic/gin"
	"wwShoppingMall/controller"
)

func CategoryRouter(engine *gin.Engine) {
	category := engine.Group("category")
	{
		category.GET("/find", controller.Category)
	}
}
