	package utils_BCCDP

	import (
		"bytes"
		"crypto/cipher"
		"crypto/des"
	)

	func PCKS5Padding(text []byte, blockSize int) []byte {
		if len(text)%blockSize == 0 { //如果明文切片 % blockSize刚好为0，不用添加
			return text
		}
		paddingSize := blockSize - len(text)%blockSize
		paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(text, paddingText...)
}
func PCKS5RemovePadding(text []byte, size int) []byte {
	//lastEle := (text[len(text)-1])
	//return text[:len(text)-lastEle]

	//
	paddingSize := text[len(text)-1]
	paddingText := bytes.Repeat([]byte{byte(paddingSize)},int(paddingSize))
	return bytes.TrimSuffix(text,paddingText)
}
func DesEnCrypt2(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originData := PCKS5Padding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	cipherData := make([]byte, len(data))
	blockMode.CryptBlocks(cipherData, originData)
	return cipherData, nil
}
/*
接口和结构体之前的联系和使用规范
接口：一组方法的签名，是一套标准
接口定义的方法只有声明(函数名，参数数量，参数类型，参数名)没有方法的实现
结构体： 结构体的字段是具体事物的特性一个描述、结构体的方法则是是这个具体事物的行为。


当结构体的一个方法的定义(方法名、方法的形参和返回值的定义的数量和顺序)和接口定义的一个方法声明完全一样时、
我们就认为 这个结构体满足了这套标准中的其中一个标准。

当结构体的多个方法的定义中有一部分方法和接口定义的所有方法声明完全一样时、
我们就认为 这个结构体满足了这套标准，它就是这个接口的一个实例化、具体这个结构体有多少个方法
接口不管，只要结构体实现了它定义的一套方法、这个结构体就是这个接口的实现，它就属于这个接口的一个实现体
这个结构体就能够称之为是这个接口、

 */