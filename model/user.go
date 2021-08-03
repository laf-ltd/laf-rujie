/*
 * @Author: laf
 * @Date: 2021-07-12 19:50:35
 * @LastEditTime: 2021-07-14 11:35:49
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\model\user.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:32; not null ; comment:用户登录名"`
	Password string `gorm:"size:64; comment:用户登录密码 "`
	Mail       string `gorm:"size:150; comment:用户邮箱"`
	Url        string `gorm:"size:150; comment:用户主页"`
	ScreenName string `gorm:"size:32;comment:用户昵称"`
	Group      string `gorm:"size:16; comment:用户组 ; default: visitor"`
}
