package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"wwShoppingMall/globle"
	"wwShoppingMall/model"
	"wwShoppingMall/utils"
)

//func LoginWX(code string) utils.Response {
//	var user model.User
//	openId, err := GetOpenId(code)
//	if err != nil {
//		return utils.ErrorMess("登录失败", nil)
//	}
//
//	user.OpenID = openId
//	if res, err := globle.UserColl.InsertOne(context.TODO(), user); err != nil {
//		return utils.ErrorMess("用户添加失败", nil)
//	} else {
//		return utils.SuccessMess("成功", res)
//	}
//	//var DBUser model.User
//	//if user.Account != "" {
//	//	//第一次登录采用账号密码
//	//	//检验账号
//	//	if err := globle.UserColl.FindOne(context.TODO(), bson.M{"account": user.Account}).Decode(&DBUser); err != nil {
//	//		return utils.ErrorMess("账号错误", err.Error())
//	//	}
//	//	//校验密码
//	//	if err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password+DBUser.Salt)); err != nil {
//	//		return utils.ErrorMess("密码错误", err.Error())
//	//	}
//	//	if DBUser.OpenId != openId {
//	//		if DBUser.OpenId == "" {
//	//			DBUser.OpenId = openId
//	//			//openId绑定
//	//			opt := bson.M{"&set": DBUser}
//	//			update := bson.M{"_id": DBUser.Id}
//	//			if _, err := globle.UserColl.UpdateOne(context.TODO(), update, opt); err != nil {
//	//				return utils.ErrorMess("绑定成功", "")
//	//			}
//	//		} else {
//	//			return utils.ErrorMess("登录失败，此账号已与其它微信账号绑定", "")
//	//		}
//	//	}
//	//} else {
//	//	//第n次登录
//	//	if err := globle.UserColl.FindOne(context.TODO(), bson.M{"openId": openId}).Decode(&DBUser); err != nil {
//	//		return utils.ErrorMess("登录失败", err.Error())
//	//	}
//	//}
//	////生成token
//	////token, err := middleware.CreateToken(DBUser)
//	//if err != nil {
//	//	return utils.ErrorMess("生成token失败", nil)
//	//}
//	//res := map[string]interface{}{
//	//	"_id":       DBUser.Id,
//	//	"account":   DBUser.Account,
//	//	"password":  DBUser.Password,
//	//	"name":      DBUser.Name,
//	//	"sex":       DBUser.Sex,
//	//	"phone":     DBUser.Phone,
//	//	"AvatarUrl": DBUser.AvatarUrl,
//	//	//"token":     token,
//	//}
//
//	res := user.SessionKey
//	return utils.SuccessMess("登录成功", res)
//}
//
//func GetOpenId(code string) (string, error) {
//	//合成url, 这里的appId和secret是在微信公众平台上获取的
//	//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
//	url := "https://api.weixin.qq.com/sns/jscode2session?" +
//		"appid=" + "wx34dbd6a627989185" + //小程序appId
//		"&secret=" + "ab6230d9bbac6a47665e1a859526b8b2" + // 小程序appSecret
//		"&js_code=" + code + //用户的code, 登录时获取
//		"&grant_type=authorization_code"
//	obj, err := http.Get(url)
//	if err != nil {
//		return "", err
//	}
//	body, _ := ioutil.ReadAll(obj.Body)
//	var res map[string]string
//	err = json.Unmarshal(body, &res)
//	if err != nil {
//		return "", err
//	}
//	return res["openid"], nil
//}

// sercret = ab6230d9bbac6a47665e1a859526b8b2

func GetOpenId(code string) (string, error) {
	//合成url, 这里的appId和secret是在微信公众平台上获取的
	//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	url := "https://api.weixin.qq.com/sns/jscode2session?" +
		"appid=" + "wx34dbd6a627989185" + //小程序appId
		"&secret=" + sercret + // 小程序appSecret
		"&js_code=" + code + //用户的code, 登录时获取
		"&grant_type=authorization_code"
	obj, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(obj.Body)
	var res map[string]string
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	return res["openid"], nil
}

func LoginWX(code string) utils.Response {
	openId, err := GetOpenId(code)
	if err != nil {
		return utils.ErrorMess("登录失败", nil)
	}
	var userwx model.UserWX
	userwx.OpenID = openId
	if res, err := globle.UserColl.InsertOne(context.TODO(), userwx); err != nil {
		return utils.ErrorMess("添加失败", nil)
	} else {
		return utils.SuccessMess("添加成功", res)
	}
}
