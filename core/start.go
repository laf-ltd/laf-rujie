/*
 * @Author: laf
 * @Date: 2021-07-12 19:57:01
 * @LastEditTime: 2021-07-12 19:57:01
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\core\start.go
 * ©低空飞行工作室(http://laf.ltd)
 */

package core

import (
	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/common/global"
	"laf.ltd/rujie/core/initialize"
)

// Start 启动WEB
func Start() {
	r := gin.Default()
	r = initialize.CollectRoute(r)

	port := global.R_VP.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run()
	}
}
