package routers

import (
	"BlockChainCertDataPorject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//默认登录页面
    beego.Router("/", &controllers.MainController{})
    //注册页面
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/register.html", &controllers.RegisterController{})
    //登录页面展示
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/login.html", &controllers.LoginController{})
    //文件上传功能
    beego.Router("/upload",&controllers.UploadFileController{})
    //实名认证
    beego.Router("/user_kyc",&controllers.KycController{})
    //查看证书
	beego.Router("/cert_detail.html",&controllers.CertDetailController{})

}
