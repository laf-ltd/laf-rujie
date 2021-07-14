/*
 * @Author: laf
 * @Date: 2021-07-14 08:25:06
 * @LastEditTime: 2021-07-14 11:57:40
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\service\user.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"laf.ltd/rujie/common"
	"laf.ltd/rujie/common/global"
	"laf.ltd/rujie/model"
)

type userService struct{}

// UserService create userService instance
var UserService = userService{}

func (u userService) Login(user model.User) (gin.H, error) {
	// 判断用户名是否为空
	if user.Name == "" {
		return nil, errors.New("登录名不能为空")
	}

	// 判断密码是否正确
	if user.Password == "" {
		return nil, errors.New("密码不能为空")
	}

	// 判断登录用户是否存在
	var userinfo model.User
	global.R_DB.Where("name=?", user.Name).First(&userinfo)
	if userinfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	// 判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(userinfo.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	// 发放token
	token, err := common.ReleaseToken(userinfo)
	if err != nil {
		return nil, err
	}

	// 返回token
	return gin.H{"token": token}, nil
}

func (u userService) Register(user model.User) (gin.H, error) {

	if len(user.Name) < 3 {
		return nil, errors.New("登录名不能少于3位")
	}

	if len(user.Mail) != 11 {
		return nil, errors.New("手机号必须为11位")
	}

	if len(user.Password) < 6 {
		return nil, errors.New("密码不能少于6位")
	}

	if isNameExist(user.Name) {
		return nil, errors.New("用户已存在")
	}

	encryptionPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("加密错误")
	}

	user.Password = string(encryptionPassword)

	global.R_DB.Create(&user)
	return gin.H{}, nil
}

func isNameExist(name string) bool {
	var user model.User
	global.R_DB.Where("name = ?", name).First(&user)
	return user.ID != 0
}
