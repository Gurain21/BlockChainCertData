package controllers

import (
	"BlockChainCertDataPorject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type KycController struct {
	beego.Controller
}

func (k *KycController) Get() {

}
func (k *KycController) Post() {
	phone := k.GetString("phone")
	var user models.User
	err := k.ParseForm(&user)
	if err != nil{
		fmt.Println(err.Error())
		k.Ctx.WriteString("信息注册失败，请稍后重试！")
		return
	}
	_,err  = user.UpdataUser()
	if err != nil {
		fmt.Println(err.Error())
		k.Ctx.WriteString("实名认证失败，请稍后重试！")
		return
	}
	k.Data["Phone"] = phone
	k.TplName = "uploadFile.html"
}
