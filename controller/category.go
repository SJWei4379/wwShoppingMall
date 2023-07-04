package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wwShoppingMall/service"
	"wwShoppingMall/utils"
)

func Category(c *gin.Context) {
	//var goods model.Goods
	//if err := c.Bind(&goods); err != nil {
	//	c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
	//}
	//c.JSON(http.StatusOK, service.FindGoods(goods))

	tag := c.Query("tag")
	if tag == "" {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", nil))
	}
	//fmt.Println(tag, "fasdfa")
	c.JSON(http.StatusOK, service.Category(tag))
}
