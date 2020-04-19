package router

import (
	"github.com/gin-gonic/gin"
	"online_shop/api"
	"online_shop/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "/static")

	userController := new(api.UserController)
	router.POST("/user/login", userController.Login)
	router.POST("/user/register", userController.Register)

	user := router.Group("/user").Use(middleware.CheckLogin())
	{
		user.GET("selfInfo", userController.SelfInfo)
	}
	return router
}
