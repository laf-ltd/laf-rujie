/*
 * @Author: laf
 * @Date: 2021-07-12 20:00:35
 * @LastEditTime: 2021-07-12 20:00:35
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\core\initialize\routes.go
 * ©低空飞行工作室(http://laf.ltd)
 */

package initialize

import (
	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/core/middleware"
	"laf.ltd/rujie/route"
)

// CollectRoute 路由配置
func CollectRoute(r *gin.Engine) *gin.Engine {

	// 配置跨域
	r.Use(middleware.Cors())

	// 公共路由 不做鉴权
	PublicGroup := r.Group("")
	{
		route.InitBaseRouter(PublicGroup)
	}

	// 私有路由需要token鉴权
	PrivateGroup := r.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		// route.CategoryRouter(PrivateGroup)
		route.UserRouter(PrivateGroup)
	}
	return r
}
