//read config file using viper
//refer https://github.com/spf13/viper
//		https://darjun.github.io/2020/01/18/godailylib/viper/

package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	APP_NAME       string
	LOG_LEVEL      string
	MYSQL_IP       string
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string
	REDIS_IP       string
	REDIS_PORT     string
)

func InitConfig() {
	//读取配置的viper工具定义
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	APP_NAME = viper.GetString("app_name")
	LOG_LEVEL = viper.GetString("log_level")

	// fmt.Println(APP_NAME)
	// fmt.Println(LOG_LEVEL)
	MYSQL_IP = viper.GetString("mysql.ip")
	MYSQL_PORT = viper.GetString("mysql.port")
	MYSQL_USER = viper.GetString("mysql.user")
	MYSQL_PASSWORD = viper.GetString("mysql.password")
	MYSQL_DATABASE = viper.GetString("mysql.database")

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))

}
