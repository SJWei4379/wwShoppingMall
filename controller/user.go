package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wwShoppingMall/service"
)

func LoginWX(c *gin.Context) {
	//var user model.User
	//if err := c.Bind(&user); err != nil {
	//	c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
	//	return
	//}
	code := c.Query("code")
	fmt.Println(code, "JHgjh")
	c.JSON(http.StatusOK, service.LoginWX(code))

	//loginData := struct {
	//	Code string `json:"code"`
	//}{}
	//
	//c.ShouldBind(&loginData)
	//
	//url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code "
	//url = fmt.Sprintf(url, appid, appSecret, loginData.Code)
	//
	//// 发起请求
	//res, _ := http.Get(url)
	//
	//// 成功后获取openId
	//wxRes := models.WXLoginRes{}
	//json.NewDecoder(res.Body).Decode(&wxRes)
	//
	//if err := dao.UserLoginDao(wxRes); err != nil {
	//	response.Err(c, "登录失败", err)
	//}
	//
	//response.Success(c, "登录成功", gin.H{
	//	"openId": wxRes.OpenID,
	//})
}
