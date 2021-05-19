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

	fmt.Println("orm得到的配置参数为", util.MYSQL_IP)
	connection := util.MYSQL_USER + ":" + util.MYSQL_PASSWORD + "@tcp(" + util.MYSQL_IP + ":" + util.MYSQL_PORT + ")/xi102collection?=utf8"
	// fmt.Println()
	engine, err = xorm.NewEngine("mysql", connection)
	if err != nil {
		fmt.Println("xorm连接失败", err)
	}

}
