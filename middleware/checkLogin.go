package middleware

import (
	"github.com/gin-gonic/gin"
	"online_shop/api"
)

func CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		uid, err := api.GetUidByHead(context)
		if err != nil || uid <= 0 {
			context.JSON(401, api.FailedResponse("not login", nil))
			context.Abort()
		}
		context.Next()
	}
}
