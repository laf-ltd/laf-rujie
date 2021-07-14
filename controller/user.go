/*
 * @Author: laf
 * @Date: 2021-07-14 08:25:06
 * @LastEditTime: 2021-07-14 11:51:34
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\controller\user.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/model"
	"laf.ltd/rujie/model/request"
	"laf.ltd/rujie/model/response"
	"laf.ltd/rujie/service"
)

func Login(ctx *gin.Context) {
	requestUser := model.User{}
	ctx.Bind(&requestUser)
	data, err := service.UserService.Login(requestUser)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, data, "登录成功")
}

func Register(ctx *gin.Context) {
	requestUser := model.User{}
	ctx.Bind(&requestUser)

	data, err := service.UserService.Register(requestUser)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, data, "注册成功")
}

func UserInfo(ctx *gin.Context) {
	user, isOk := ctx.Get("user")
	if !isOk {
		response.Error(ctx, http.StatusBadRequest, "获取用户信息失败")
		return
	}
	response.Success(ctx, gin.H{"user": request.ToUserRequest(user.(model.User))}, "获取用户信息成功")
}
