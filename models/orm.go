package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xi102/collection/util"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitDatabase() {
	var err error

	fmt.Println("orm得到的配置参数为", util.MYSQL_IP)
	connection := util.MYSQL_USER + ":" + util.MYSQL_PASSWORD + "@tcp(" + util.MYSQL_IP + ":" + util.MYSQL_PORT + ")/xi102collection?charset=utf8"
	// fmt.Println()
	// engine, err = xorm.NewEngine("mysql", connection)
	// if err != nil {
	// 	fmt.Println("xorm连接失败", err)
	// }
	// sql := "select * from userinfo"
	// results, err := engine.Query(sql)
	// for i := range results {
	// 	fmt.Printf(results[i][0])
	// }
	// fmt.Printf(results)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mysql 连接成功")

	rows, err := db.Query("select * from user where id = 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//打印执行结果

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}

			fmt.Println(col, v)
		}
	}

	defer db.Close()

}
