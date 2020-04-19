package api

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"online_shop/service/jwtx"
)

const (
	ParamError = "ParamError"
	Success    = "Success"
	Failed     = "Failed"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func SuccessResponse(msg string, data interface{}) Response {
	return Response{
		Code: 200,
		Data: data,
		Msg:  msg,
	}
}

func FailedResponse(msg string, data interface{}) Response {
	return Response{
		Code: 500,
		Data: data,
		Msg:  msg,
	}
}

func GetResponse(code int, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

func GetUidByHead(context *gin.Context) (int, error) {
	tokenString := context.Request.Header.Get("Authorization")
	if tokenString == "" {
		return 0, errors.New(ParamError)
	}
	claims, err := jwtx.ParseToken(tokenString)
	if err != nil {
		logs.Error(err)
		return 0, errors.New(ParamError)
	}
	uid := int(claims["uid"].(float64))
	return uid, nil
}
