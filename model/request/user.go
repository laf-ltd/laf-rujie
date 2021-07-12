/*
 * @Author: laf
 * @Date: 2021-07-12 20:25:45
 * @LastEditTime: 2021-07-12 19:47:37
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\model\request\user.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package request

import "laf.ltd/rujie/model"

// LoginRequest LoginRequest
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// RegisterRequest RegisterRequest
type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// UserRequest UserRequest
type UserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// ToUserRequest ToUserRequest
func ToUserRequest(user model.User) UserRequest {
	return UserRequest{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
