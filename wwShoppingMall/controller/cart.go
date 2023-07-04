package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"wwShoppingMall/service"
	"wwShoppingMall/utils"
)

func AddCart(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	image := c.Query("image")

	currentID, err := primitive.ObjectIDFromHex(c.Query("currentID"))
	if title == "" || price == "" || image == "" || err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", nil))
		return
	}
	c.JSON(http.StatusOK, service.AddCart(title, price, image, currentID))
}

func GetCart(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetCart())
}

func DelCart(c *gin.Context) {
	if _id, err := primitive.ObjectIDFromHex(c.Query("_id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.DelCart(_id))
	}
}
