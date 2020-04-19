package api

import (
	"github.com/gin-gonic/gin"
	"online_shop/model"
)

type UserController struct {
}

func (t *UserController) Login(c *gin.Context) {
	params := make(map[string]string)
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, FailedResponse(err.Error(), nil))
		return
	}
	mobile := params["mobile"]
	password := params["password"]
	token, err := new(model.User).Login(mobile, password)
	if err != nil {
		c.JSON(400, FailedResponse(err.Error(), nil))
		return
	}
	ret := make(map[string]string)
	ret["token"] = token
	c.JSON(200, SuccessResponse("success login", ret))
}

func (t *UserController) Register(c *gin.Context) {
	params := make(map[string]string)
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(500, FailedResponse(err.Error(), nil))
		return
	}
	mobile := params["mobile"]
	password := params["password"]
	token, err := new(model.User).Register(mobile, password)
	if err != nil {
		c.JSON(400, FailedResponse(err.Error(), nil))
		return
	}
	ret := make(map[string]string)
	ret["token"] = token
	c.JSON(200, SuccessResponse("success register", ret))
}

func (t *UserController) SelfInfo(c *gin.Context) {
	uid, _ := GetUidByHead(c)
	user, err := new(model.User).Info(uid)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, SuccessResponse("获取当前用户信息", *user))

}
