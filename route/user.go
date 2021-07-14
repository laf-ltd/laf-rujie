/*
 * @Author: laf
 * @Date: 2021-07-14 08:25:06
 * @LastEditTime: 2021-07-14 11:52:22
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\route\user.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package route

import (
	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/controller"
)

// UserRouter 用户信息接口 路由
func UserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	BaseRouter := Router.Group("user")
	{
		BaseRouter.GET("info", controller.UserInfo)
	}
	return BaseRouter
}
