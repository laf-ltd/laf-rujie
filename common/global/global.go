/*
 * @Author: laf
 * @Date: 2021-07-12 19:30:35
 * @LastEditTime: 2021-07-12 19:30:35
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\core\global.go
 * ©低空飞行工作室(http://laf.ltd)
 */
package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	R_VP *viper.Viper
	R_DB *gorm.DB
)
