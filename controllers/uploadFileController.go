package controllers

import (
	"BlockChainCertDataPorject/block_chain"
	"BlockChainCertDataPorject/models"
	"BlockChainCertDataPorject/utils_BCCDP"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type UploadFileController struct {
	beego.Controller
}

func (u *UploadFileController) Get() {
	u.TplName = "uploadFile.html"
}

func (u *UploadFileController) Post() {
	//1、先看用户又没有实名认证：如果没有 跳转到实名认证页面
	var user1 models.User
	err := u.ParseForm(&user1)
	fmt.Println("用户的电话号码为:",user1.Phone)

	user, err := user1.QueryUserByPhone()
	if err != nil {
		fmt.Println("查询用户信息错误", err.Error())
		u.Ctx.WriteString("查询用户信息错误")
		return
	}
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Card) == "" || strings.TrimSpace(user.Sex) == "" {
		u.Data["Phone"] = user1.Phone
		u.TplName = "user_kyc.html"
		return
	}
	//2、实名认证后，判断用户提交的文件的类型,文件大小
	title := u.GetString("upload_title")
	file, fileHeader, err := u.GetFile("file")
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("上传文件遇到错误，请稍后重试！")
		return
	}

	defer file.Close()
	isJpg := strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".jpg")
	isPng := strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".png")
	if !isJpg && !isPng {
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
		return
	}
	if fileHeader.Size/1024 > 1000 {
		u.Ctx.WriteString("抱歉，文件大小不符合, 请上传符合大小的文件")
		return
	}
	//3、用户提交的文件符合我们的要求时，保存文件到本地，
	saveDir := "static/upload"
	utils_BCCDP.OpenDir(saveDir) //打开要保存文件的本地目录
	//文件名： 文件路径 + 文件名 + "." + 文件扩展名Filename
	saveFileName := saveDir + "/" + fileHeader.Filename
	fmt.Println("要保存的文件名", saveFileName)
	utils_BCCDP.SaveFile(saveFileName, file)
	//4、把文件的信息保存到数据库中，再上到区块链
	//①文件的md5值和sha256值先计算出来
	file, _ = fileHeader.Open()
	fileMd5Hash, _ := utils_BCCDP.MD5HashReader(file)
	file, _ = fileHeader.Open()
	fileBytes,err := ioutil.ReadAll(file)
	fileSha256HashBytes:= utils_BCCDP.SHA256HashByte(fileBytes)
	file, _ = fileHeader.Open()
	fileMd5HashBytes:= utils_BCCDP.MD5HashByte(fileBytes)

	//②保存到数据库中
	record := models.UploadRecord{
		UserId:    user.Id,
		FileName:  fileHeader.Filename,
		FileSize:  fileHeader.Size,
		FileCert:  fileMd5Hash,
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}
	_, err = record.SaveRecord()
	if err != nil {
		fmt.Println("保存认证记录:", err.Error())
		u.Ctx.WriteString("抱歉，电子数据认证保存失败，请稍后再试!")
		return
	}
	//③保存到区块链
	certRecord := models.CertRecord{
		CertId:   fileMd5HashBytes,
		CertHash: fileSha256HashBytes,
		CertName: user.Name,
		Phone:    user.Phone,
		CertCard: user.Card,
		FileName: fileHeader.Filename,
		FileSize: fileHeader.Size,
		CertTime: time.Now().Unix(),
	}
	certRecordBytes, err := certRecord.Serialize() //序列化
	block, err := block_chain.CHAIN.SaveData(certRecordBytes)
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉,数据上链失败" + err.Error())
		return
	}
	fmt.Println("恭喜,数据保存到区块链成功!区块的高度为:", block.Height)

	records, err := models.QueryRecordsByUserId(user.Id)
	if err != nil {
		fmt.Println("获取数据列表:", err.Error())
		u.Ctx.WriteString("抱歉, 获取电子数据列表失败, 请重新尝试!")
		return
	}
	u.Data["Phone"] = user1.Phone
	u.Data["Records"] = records
	u.TplName = "list_record.html"
}
