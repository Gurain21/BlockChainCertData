package controllers

import (
	"BlockChainCertDataPorject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

//处理登录提交的表单
func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，登录遇到错误，请稍后再试！")
		return
	}
	u,err := user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，登录遇到错误，请检查密码后重试！")
		return
	}
	l.Data["Phone"] = u.Phone
	l.TplName = "uploadFile.html"

}
