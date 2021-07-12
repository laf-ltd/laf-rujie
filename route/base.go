package route

import (
	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/controller"
)

// InitBaseRouter 初始化基础路由
func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	BaseRouter := Router.Group("auth")
	{
		UserRouter := controller.UserController
		BaseRouter.POST("login", UserRouter.Login)
		BaseRouter.POST("register", UserRouter.Register)
	}
	return BaseRouter
}
