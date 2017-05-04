package config

import (
	"tools/file"
	"tools/logger"
)

const (
	//配置文件名称
	configFile = "config.conf"
)

//检查或初始化配置文件
func init() {

	if file.Exist(configFile) {
		logger.Debug("config.conf存在，开始读取配置……")
	} else {
		logger.Info("创建配置文件config.conf！")
	}
}
