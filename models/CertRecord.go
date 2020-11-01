package models

import (
	"bytes"
	"encoding/gob"
)

//该struct用于定义链上数据保存的信息
type CertRecord struct {
	CertId         []byte //文件的md5 哈希值
	CertHash       []byte //文件的sha256哈希值
	CertName       string //保存人姓名
	Phone          string //保存人电话
	CertCard       string //保存人身份证
	FileName       string //文件名
	FileSize       int64  //文件大小
	CertTime       int64  //文件保存时间
	CertIdFormat   string //仅用于md5 []byte格式化输出
	CertHashFormat string //仅用于sha256 []byte格式化输出
	CertTimeFormat string //仅用于时间格式化输出
}

func (c CertRecord) Serialize() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(), err
}
func DeSerializeCertRecord(data []byte) (*CertRecord, error) {

	var certRecord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	return certRecord, err
}
