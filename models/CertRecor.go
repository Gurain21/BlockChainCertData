package models

import (
	"bytes"
	"encoding/gob"
)

//该struct用于定义链上数据保存的信息
type CertRecord struct {
	CertId []byte
	CertHash []byte
	CertName string
	Phone string
	CertCard string
	FileName string
	FileSize int64
	CertTime int64
}

func (c CertRecord)Serialize()([]byte,error){
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(),err
}
func DeSerializeCertRecord(data []byte)(*CertRecord,error)  {

	var certRecord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	return  certRecord,err
}
