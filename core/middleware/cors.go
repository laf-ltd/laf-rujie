/*
 * @Author: laf
 * @Date: 2021-07-12 20:16:41
 * @LastEditTime: 2021-07-12 19:47:25
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\core\middleware\cors.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
