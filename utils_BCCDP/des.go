package utils_BCCDP

import (
	"crypto/cipher"
	"crypto/des"
)
func TripleEnCrypt(text,key []byte)[]byte{
	//text分成若干个 ke y[]byte的长度大小,不足则补齐
	block ,_ := des.NewTripleDESCipher(key)
	//选择des加密的模式 ： CBC：密码分组链接模式 cipher block chaining 密码 块 链接
	enDesBytes := make([]byte,len(key))
	blockMode := cipher.NewCBCEncrypter(block,key[:block.BlockSize()])
	blockMode.CryptBlocks(enDesBytes,PCKS5Padding(text,block.BlockSize()))
	return  enDesBytes
}
func DesEnCrypt(text,key []byte)[]byte{
	//text分成若干个 ke y[]byte的长度大小,不足则补齐
	block ,_ := des.NewCipher(key)
	//选择des加密的模式 ： CBC：密码分组链接模式 cipher block chaining 密码 块 链接
	enDesBytes := make([]byte,len(key))
	blockMode := cipher.NewCBCEncrypter(block,enDesBytes)
	blockMode.CryptBlocks(enDesBytes,PCKS5Padding(text,block.BlockSize()))
	return  enDesBytes
}
func DesDeCrypt(text,key []byte)[]byte  {
	block,_ := des.NewCipher(key)
	deDesBytes := make([]byte,len(key))
	blockMode := cipher.NewCBCDecrypter(block,deDesBytes)
	blockMode.CryptBlocks(deDesBytes,text)
	return PCKS5RemovePadding(deDesBytes,block.BlockSize())
}
