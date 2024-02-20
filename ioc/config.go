package ioc

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("dev")      //不包含文件扩展名
	viper.SetConfigType("yaml")     //扩展名
	viper.AddConfigPath("./config") //会是当前目录的工作目录下
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		//panic(fmt.Errorf("fatal error config file: %w", err))
		panic(fmt.Errorf("配置文件读取失败: %w", err))
	}
}

// InitViperV3Remote 远程配置放在etcd里面
func InitViperV3Remote() {
	viper.SetConfigType("yaml")
	//切换
	//viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
	err := viper.AddRemoteProvider("etcd3",
		"127.0.0.1:12379", "/webook/config")
	if err != nil {
		panic(err)
	}
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
}
