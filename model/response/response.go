/*
 * @Author: laf
 * @Date: 2021-07-12 20:24:53
 * @LastEditTime: 2021-07-12 19:47:46
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\model\response\response.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response Response
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Result Result
func Result(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, Response{
		code,
		data,
		msg,
	})
}

// Success Success
func Success(ctx *gin.Context, data gin.H, msg string) {
	Result(ctx, http.StatusOK, 200, data, msg)
}

// Ok ok
func Ok(ctx *gin.Context, data gin.H) {
	Result(ctx, http.StatusOK, 200, data, "操作成功")
}

// Error Error
func Error(ctx *gin.Context, status int, msg string) {
	Result(ctx, http.StatusOK, status, gin.H{}, msg)
}

// Fail Fail
func Fail(ctx *gin.Context) {
	Result(ctx, http.StatusOK, 400, gin.H{}, "操作失败")
}
