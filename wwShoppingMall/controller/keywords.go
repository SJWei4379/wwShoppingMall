package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wwShoppingMall/service"
)

func GetKeywords(c *gin.Context) {
	//var banner model.Banner
	//c.Query("banner")
	c.JSON(http.StatusOK, service.GetKeywords())
}
