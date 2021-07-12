package controller

import (
	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/model"
	"laf.ltd/rujie/service"
	"net/http"
)

// UserController UserController
type userController struct {
}

var UserController = userController{}

func (u userController) Login(ctx *gin.Context) {
	requestUser := model.User{}
	ctx.Bind(&requestUser)
	data, err := service.UserService.Login(requestUser)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, data, "登录成功")
}

func (u userController) Register(ctx *gin.Context) {
	requestUser := model.User{}
	ctx.Bind(&requestUser)

	data, err := service.UserService.Register(requestUser)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, data, "注册成功")
}

func (u userController) UserInfo(ctx *gin.Context) {
	user, isOk := ctx.Get("user")
	if !isOk {
		response.Error(ctx, http.StatusBadRequest, "获取用户信息失败")
		return
	}
	response.Success(ctx, gin.H{"user": request.ToUserRequest(user.(model.User))}, "获取用户信息成功")
}
