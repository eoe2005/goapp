package db

import (
	"database/sql"
	"fmt"
)

func showcreateTable(dbcon *sql.DB, tbName string) {
	rows, e := dbcon.Query("SHOW CREATE TABLE " + tbName)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer rows.Close()
	var name, sql string
	if rows.Next() {
		rows.Scan(&name, &sql)
	}
	buildGo(sql)
}
