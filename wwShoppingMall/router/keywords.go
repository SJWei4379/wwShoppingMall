package router

import (
	"github.com/gin-gonic/gin"
	"wwShoppingMall/controller"
)

func KeywordsRouter(engine *gin.Engine) {
	keywords := engine.Group("keywords")
	{
		keywords.GET("/get", controller.GetKeywords)
	}
}
