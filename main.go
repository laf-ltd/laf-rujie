/*
 * @Author: laf
 * @Date: 2021-07-12 20:22:27
 * @LastEditTime: 2021-07-12 20:28:13
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\main.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package main

import (
	"laf.ltd/rujie/common/global"
	"laf.ltd/rujie/core"
	"laf.ltd/rujie/core/initialize"
)

func main() {
	// 初始化Viper
	global.R_VP = core.Viper()
	// gorm链接数据库
	global.R_DB = initialize.InitDB()
	// 初始化表
	initialize.AutoMigrate(global.R_DB)

	// // 程序结束前关闭数据库链接
	db, _ := global.R_DB.DB()
	defer db.Close()

	// 启动 gin
	core.Start()

}
