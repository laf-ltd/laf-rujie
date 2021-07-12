/*
 * @Author: laf
 * @Date: 2021-07-12 19:32:26
 * @LastEditTime: 2021-07-12 19:32:27
 * @LastEditors: laf
 * @Description:
 * @FilePath: \laf-rujie\core\viper.go
 * ©低空飞行工作室(http://laf.ltd)
 */

package core

import (
	"os"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	workDir, _ := os.Getwd()
	v.SetConfigName("application")
	v.SetConfigType("yml")
	v.AddConfigPath(workDir + "/config")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return v
}
