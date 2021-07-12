package route

import (
	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/controller"
)

// UserRouter 用户信息接口 路由
func UserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	BaseRouter := Router.Group("user")
	{
		UserRouter := controller.UserController
		BaseRouter.GET("info", UserRouter.UserInfo)
	}
	return BaseRouter
}
