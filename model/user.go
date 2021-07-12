/*
 * @Author: laf
 * @Date: 2021-07-12 19:50:35
 * @LastEditTime: 2021-07-12 19:50:35
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\model\user.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(20);not null"`
	Phone    string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
