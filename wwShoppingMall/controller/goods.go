package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"wwShoppingMall/service"
	"wwShoppingMall/utils"
)

func GetGoods(c *gin.Context) {
	//pageSize := c.Query("pageSize")
	currPage := c.Query("page")
	if currPage == "" {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", nil))
	}
	c.JSON(http.StatusOK, service.GetGoods(currPage))
}

func FindGoods(c *gin.Context) {
	//var goods model.Goods
	//if err := c.Bind(&goods); err != nil {
	//	c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
	//}
	//c.JSON(http.StatusOK, service.FindGoods(goods))

	search := c.Query("search")
	if search == "" {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", nil))
	}
	c.JSON(http.StatusOK, service.FindGoods(search))
}

func BuyFindGoods(c *gin.Context) {
	if _id, err := primitive.ObjectIDFromHex(c.Query("id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.BuyFindGoods(_id))
	}
}
