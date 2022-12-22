package main

import (
	_ "embed"

	"github.com/eoe2005/goapp/sql2gorm/db"
)

// //go:embed tt.sql
// var sqlData string

func main() {
	db.GetDb()
	// db.Test(sqlData)
}
