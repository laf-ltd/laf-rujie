/*
 * @Author: laf
 * @Date: 2021-07-12 19:42:07
 * @LastEditTime: 2021-07-14 09:04:51
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\core\initialize\gorm.go
 * ©低空飞行工作室(http://laf.ltd)
 */

package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"laf.ltd/rujie/common/global"
	"laf.ltd/rujie/model"
)

// InitDB 初始化数据
func InitDB() *gorm.DB {
	dsn := global.R_VP.GetString("mysql.dsn")

	// https://github.com/go-sql-driver/mysql
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	return db
}

// AutoMigrate 自动迁移（初始化表）
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		panic(err)
	}
}
