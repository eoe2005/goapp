package db

import (
	"database/sql"
	"fmt"

	"github.com/eoe2005/goapp/mysql2gorm/console"
	_ "github.com/go-sql-driver/mysql"
)

func GetDb() {
	host := console.ReadDefault("请输入服务器地址[默认: localhost] -> ", "127.0.0.1")
	port := console.ReadDefault("请输入端口[默认: 3306] -> ", "3306")
	user := console.ReadDefault("请输入账号[默认: root] -> ", "root")
	pass := console.ReadDefault("请输入密码 -> ", "")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port))
	if err != nil {
		fmt.Printf("链接数据库错误 ： %s\n\n", err.Error())
		GetDb()
		return
	}
	defer db.Close()
	selectDB(db)
	selectTable(db)
}
