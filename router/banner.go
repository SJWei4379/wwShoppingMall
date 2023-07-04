package router

import (
	"github.com/gin-gonic/gin"
	"wwShoppingMall/controller"
)

func BannerRouter(engine *gin.Engine) {
	banner := engine.Group("banner")
	{
		banner.GET("/find", controller.BannerFind)
	}
}
