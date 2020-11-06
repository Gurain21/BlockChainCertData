package utils_BCCDP

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)
/*
//dy:dynamic  动态的  batch 批量,批次

accessKeyId:LTAI4FydRQswZLg94JhLGRLp

accessKeySercet: 8Stma2gzzZaGYGo0bQckBMziHO3FWx


 */
const SMS_TLP_REGISTER = "SMS_205393604" //注册业务的短信模板
const SMS_TLP_LOGIN = "SMS_205398654"    //用户登录的短信模板
const SMS_TLP_KYC = ""      //实名认证的短信模板

//该函数用于向发送一条短信 phone 手机号  code 验证码   templateType 模板类型
func SendSms(phone string,code string,templateType string) (*SmsResult,error) {
	config := beego.AppConfig
	sms_access_key := config.String("sms_access_key")
	sms_access_secret  :=config.String("sms_access_secret")
	regiold :=config.String("regiold")
	client, err := dysmsapi.NewClientWithAccessKey(regiold, sms_access_key, sms_access_secret)
	if err != nil{
		return nil,err
	}
//创建一个发送短信服务的请求
	requset:= dysmsapi.CreateSendSmsRequest()
	//指定要发送个的目标手机号
	requset.PhoneNumbers = phone
	//指定签名信息
	requset.SignName ="线上餐厅"
	//指定短信模板  编号
	requset.TemplateCode = templateType
	smsCode := SmsCode{Code:code}
	smsBytes,_ := json.Marshal(smsCode)
	requset.TemplateParam = string(smsBytes)

	response,err := client.SendSms(requset)
	if err != nil {
		return nil,err
	}
	//Biz business 商业 | 业务
	smsResult := &SmsResult{
		BizId:     response.BizId,
		Code:      response.Code,
		Message:   response.Message,
		RequestId: response.RequestId,
	}
	return smsResult,nil
}


type SmsCode struct {
	Code string `json:"code"`
}
type SmsResult struct {
	BizId string
	Code string
	Message string
	RequestId string
}

func GenRandCode(width int)string  {
	//numeric :=[10]byte{0,1,2,3,4,5,6,7,8,9}
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i:=0;i<width ;i++  {
	fmt.Fprintf(&sb,"%d",byte(rand.Intn(10)))
	}
return  sb.String()
}