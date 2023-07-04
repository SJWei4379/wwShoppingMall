package router

import (
	"github.com/gin-gonic/gin"
	"wwShoppingMall/controller"
	"wwShoppingMall/middleware"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()                         //请求默认的gin初始化启动
	engine.Use(middleware.Log(), middleware.Cors()) //日志   跨域

	engine.POST("loginWX", controller.LoginWX) //微信登录

	UserRouter(engine)
	BannerRouter(engine)
	GoodsRouter(engine)
	KeywordsRouter(engine)
	CartRouter(engine)
	CategoryRouter(engine)

	return engine
}
