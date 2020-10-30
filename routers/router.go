package routers

import (
	"BlockChainCertDataPorject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/re", &controllers.MainController{})
}
