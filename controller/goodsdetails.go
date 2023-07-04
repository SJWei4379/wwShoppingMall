package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"wwShoppingMall/service"
	"wwShoppingMall/utils"
)

func FindGoodsDetails(c *gin.Context) {

	if _id, err := primitive.ObjectIDFromHex(c.Query("id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.FindGoodsDetails(_id))
	}
}
