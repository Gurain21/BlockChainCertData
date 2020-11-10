package utils_BCCDP

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func EnDes(text,key []byte)[]byte{
	//text分成若干个 ke y[]byte的长度大小,不足则补齐
	block ,_ := des.NewCipher(key)
	//选择des加密的模式 ： CBC：密码分组链接模式 cipher block chaining 密码 块 链接
	enDesBytes := make([]byte,len(key))
	blockMode := cipher.NewCBCEncrypter(block,enDesBytes)
	blockMode.CryptBlocks(enDesBytes,EndAddPadding(text,block.BlockSize()))
	return  enDesBytes
}
func DeDes(text,key []byte)[]byte  {
	block,_ := des.NewCipher(key)
	deDesBytes := make([]byte,len(key))
	blockMode := cipher.NewCBCDecrypter(block,deDesBytes)
	blockMode.CryptBlocks(deDesBytes,text)
	return EndRemovePadding(deDesBytes)
}
func EndAddPadding(text []byte,blockSize int)[]byte{
	paddingSize :=blockSize - len(text)%blockSize
	paddingText :=bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	return append(text, paddingText...)
}
func EndRemovePadding(text []byte)[]byte  {
	paddingSize := text[len(text)-1]
	paddingText := bytes.Repeat([]byte{byte(paddingSize)},int(paddingSize))
	return bytes.TrimSuffix(text,paddingText)
}