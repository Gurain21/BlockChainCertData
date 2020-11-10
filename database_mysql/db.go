package database_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var DB_BCCDP *sql.DB
func OpenDB() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/block_chain_cert_db?charset=utf8")

	if err !=nil {
		panic(err.Error())
	}
	DB_BCCDP = db
	db.Driver().Open()
}
