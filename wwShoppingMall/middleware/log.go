package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
	"wwShoppingMall/globle"
	"wwShoppingMall/model"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path      //请求路径
		query := c.Request.URL.RawQuery //query参数
		c.Next()
		//ResBody := c.Request.Body
		//bodyByte, err := ioutil.ReadAll(ResBody)
		//if err != nil {
		//	fmt.Println(err.Error())
		//}
		var body interface{}
		//fmt.Println("测试JSON错误")
		//err = json.Unmarshal(bodyByte, &body)
		//if err != nil {
		//	fmt.Println(err)
		//	body = string(bodyByte)
		//}
		//fmt.Println(body)
		cost := time.Since(start) //访问时间
		user, _ := c.Get("user")
		log := model.Log{
			User:      user,
			Path:      path,
			Method:    c.Request.Method,
			Status:    c.Writer.Status(),
			Query:     query,
			Body:      body,
			Ip:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Errors:    c.Errors.ByType(gin.ErrorTypePrivate).String(),
			Cost:      cost.String(),
		}
		if log.Status == 204 {
			return
		}
		_, _ = globle.LogColl.InsertOne(context.Background(), log)
		//_,_ = globle.DB.Exec("insert into log (id,username,account,password,age,sex,salt) values (?,?,?,?,?,?,?)", user.Id, user.Username, user.Account, user.Password, user.Age, user.Sex, user.Salt)
	}
}
