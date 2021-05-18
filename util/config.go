//read config file using viper
//refer https://github.com/spf13/viper
//		https://darjun.github.io/2020/01/18/godailylib/viper/

package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	host     string
	port     int
	user     string
	password string
	database string
}
type DBConfig struct {
	MySQL MySQLConfig
}

func Config() DBConfig {
	//读取配置的viper工具定义
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf")

	//读取配置失败处理
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	//读取
	var c DBConfig
	viper.Unmarshal(&c)
	fmt.Println(c.MySQL)

	return c
}
