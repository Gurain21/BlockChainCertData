package controllers

import (
	"BlockChainCertDataPorject/block_chain"
	"BlockChainCertDataPorject/models"
	"BlockChainCertDataPorject/utils_BCCDP"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController)Get()  {
	//1、解析和接收前端页面传递的数据cert_id
	cert_id := c.GetString("cert_id")

	//2、到区块链上查询区块数据
	block, err := block_chain.CHAIN.QueryBlockByCertId(cert_id)
	if err != nil {
		c.Ctx.WriteString("抱歉，查询链上数据遇到错误，请重试！")
		return
	}
	if block == nil { //遍历整条区块链，但是未查询到数据
		c.Ctx.WriteString("抱歉，未查询到链上数据")
		return
	}
	fmt.Println("查询到区块的高度：", block.Height)

	//反序列化
	certRecord, err := models.DeSerializeCertRecord(block.Data)
	fmt.Println(certRecord)
	//BUG 显示为十进制
	/*

	65336230633434323938666331633134396166626634633839393666623932343237616534316534363439623933346361343935393931623738353262383535
	6366643964653465373461333566663732333339393966363962313335373536

	?
	 */
	certRecord.CertIdFormat =strings.ToTitle( hex.EncodeToString(certRecord.CertId))
	certRecord.CertHashFormat = hex.EncodeToString(certRecord.CertHash)
	certRecord.CertTimeFormat = utils_BCCDP.TimeFormat(certRecord.CertTime,utils_BCCDP.TIME_FORMAT_ONE)
//fmt.Println(hex.EncodeToString(certRecord.CertHash))
	//结构体
	c.Data["CertRecord"] = certRecord
	c.Data["BlockHash"] = hex.EncodeToString(block.Hash)

	//3、跳转证书详情页面
	c.TplName = "cert_detail.html"
}
func (c *CertDetailController)Post()  {

}