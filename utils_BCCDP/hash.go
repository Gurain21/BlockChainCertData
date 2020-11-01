package utils_BCCDP

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)
//对字符串 data进行MD5 哈希 返回data的md5哈希值
func MD5HashString(data string)string  {

	hashMd5 := md5.New()
	hashMd5.Write([]byte(data))
	bytes := hashMd5.Sum(nil)
	return hex.EncodeToString(bytes)
}
//读取 io流中的数据，并对数据进行哈希计算，返回md5哈希值和error
func MD5HashReader(reader io.Reader)(string,error)  {
	readerBytes,err := ioutil.ReadAll(reader)
	if err !=nil {
		return "",err
	}
	bytes :=MD5HashByte(readerBytes)
	return hex.EncodeToString(bytes),nil
}
func SHA256HashReader(reader io.Reader) (string, error) {
	sha256Hash := sha256.New()
	readerBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
//对 []byte 进行md5哈希，返回md5哈希值和error
func MD5HashByte(data []byte) ([]byte) {
	hashMd5 := md5.New()
	hashMd5.Write(data)
	return hashMd5.Sum(nil)
}
// 对[]byte 进行sha256哈希 返回sha256哈希后的byte
func SHA256HashByte(data []byte)([]byte)  {
	hashSHA256 := sha256.New()
	hashSHA256.Write(data)
	return hashSHA256.Sum(nil)
}