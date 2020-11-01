package controllers

import (
	"BlockChainCertDataPorject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}
func (r *RegisterController) Post() {
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("注册遇到错误，请稍后重试！")
		return
	}
	fmt.Println(user)
	_, err = user.AddUser()
	if err != nil {
		fmt.Println(err)
		r.Ctx.WriteString("注册失败，请稍后重试！")
	}

	r.TplName = "login.html"
}
