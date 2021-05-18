package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xi102/collection/util"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitDatabase() {
	var err error
	config := util.Config()

	//用法：config.MySQL.host
	connection := config.MySQL.user + ":" + config.MySQL.password + "@" + config.MySQL.host + "?charset=utf8"
	fmt.Println("connection")
	engine, err = xorm.NewEngine("mysql", connection)

}
