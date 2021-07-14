/*
 * @Author: laf
 * @Date: 2021-07-14 08:25:06
 * @LastEditTime: 2021-07-14 11:52:01
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\route\base.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package route

import (
	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/controller"
)

// InitBaseRouter 初始化基础路由
func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	BaseRouter := Router.Group("auth")
	{
		BaseRouter.POST("login", controller.Login)
		BaseRouter.POST("register", controller.Register)
	}
	return BaseRouter
}
