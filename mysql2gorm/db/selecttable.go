package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/eoe2005/goapp/mysql2gorm/console"
)

func selectTable(dbcon *sql.DB) {
	rows, e := dbcon.Query("SHOW TABLES")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer rows.Close()
	tbs := []string{}
	fmt.Println("数据库列表：")
	for rows.Next() {
		var name string
		rows.Scan(&name)
		tbs = append(tbs, name)
		fmt.Printf("  %s\n", name)
	}

	tbname := ""
	isOk := true
	for isOk {
		tbname = console.Read(fmt.Sprintf("选择表[ 默认全部 ] -> "))
		if tbname == "" {
			isOk = false
			break
		}

		for _, name := range tbs {
			if name == tbname {
				isOk = false
				break
			}
		}
	}
	ts := []string{}
	if tbname == "" {
		ts = append(ts, tbs...)
	} else {
		ts = strings.Split(tbname, ",")
	}
	for _, t := range ts {
		showcreateTable(dbcon, t)
	}
}
