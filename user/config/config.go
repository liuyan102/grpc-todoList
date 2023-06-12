package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd()                 // 工作目录路径
	viper.SetConfigName("config")            // 配置文件的文件名
	viper.SetConfigType("yml")               // 配置文件的后缀
	viper.AddConfigPath(workDir + "/config") // 获取到配置文件的路径
	err := viper.ReadInConfig()              // 读取配置文件
	if err != nil {
		return
	}
}
