package main

import (
	"BlockChainCertDataPorject/block_chain"
	"BlockChainCertDataPorject/database_mysql"
	_ "BlockChainCertDataPorject/routers"
	"BlockChainCertDataPorject/utils_BCCDP"
	"github.com/astaxie/beego"
)

func main() {
	//准备好一条区块链
	block_chain.NewBlockChain()
	//打开数据库
	database_mysql.OpenDB()
	defer func() {
		err := database_mysql.DB_BCCDP.Close()
		utils_BCCDP.CheckErrore(err,"关闭数据库遇到错误,h_h")
	}()



	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}

