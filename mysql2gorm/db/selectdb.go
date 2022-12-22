package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/eoe2005/goapp/sql2gorm/console"
)

func selectDB(dbcon *sql.DB) {
	rows, e := dbcon.Query("SHOW DATABASES")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer rows.Close()
	dbs := []string{}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		dbs = append(dbs, name)
	}

	dbname := ""
	isOk := true
	for isOk {
		dbname = console.Read(fmt.Sprintf("选择数据[%s] -> ", strings.Join(dbs, ",")))
		if dbname == "" {
			continue
		}

		for _, name := range dbs {
			if name == dbname {
				isOk = false
				break
			}
		}
	}
	dbcon.Exec("USE " + dbname)
}
