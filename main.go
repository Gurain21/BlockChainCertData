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
		utils_BCCDP.CheckErrore(err, "关闭数据库遇到错误,h_h")
	}()
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
	beego.Run(":8088")
	/*  1、连接远程服务器   finlshell工具或其它工具
		2、编译我们的程序   go build main.go  公网8088端口*	main.exe 把该程序上传到我们的云服务器上
	  	3、在云服务器上通过cmd命令执行这个main.exe ,   设置云服务器的端口 入口和出口(服务器有防火墙等安全设置,无法直接访问).    (项目端口号)   授权对象:所有人      任何人都能访问
		 												打开防火墙的端口
	 */


}
