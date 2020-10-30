package main

import (
	"BlockChainCertDataPorject/database_mysql"
	_ "BlockChainCertDataPorject/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	database_mysql.OpenDB()
	fmt.Println(database_mysql.DB_BCCDP)
	beego.Run()
}

