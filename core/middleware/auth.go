/*
 * @Author: laf
 * @Date: 2021-07-12 20:02:54
 * @LastEditTime: 2021-07-12 19:46:49
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\core\middleware\auth.go
 * ©低空飞行工作室(http://laf.ltd)
 */

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"laf.ltd/rujie/common"
	"laf.ltd/rujie/common/global"
	"laf.ltd/rujie/model"
)

// JWTAuth jwt鉴权
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后获取Claim中的userId
		userID := claims.UserID

		var user model.User
		global.R_DB.First(&user, userID)

		//用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
