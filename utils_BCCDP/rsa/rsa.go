package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"math/big"
	"os"
)
/**RSA算法加密步骤：
*   1、生成一个私钥，公钥为私钥这个结构体的一个字段名   ：
①：var bits int  ②：flag.IntVar(&bits,"name",1024,"usage")③：key, err :=rsa.GenerateKey(rand.Reader,bits)
*	2、将这个私钥和公钥写入到一个后缀为pem的文件中(持久化存储密钥)：
①：priStream := x509.MarshalPKCS1PrivateKey(key) ②：block1 := &pem.Block{ Type: " RSA private key ",Bytes:  pubStream,}
③：file,err := os.Create("filename")  ④：err :=  pem.Encode(file,blocks)
   3、rsa加密：公钥加密，私钥解密：
	①:cryptData,err := rsa.EncryptPKCS1v15(rand.Reader, &key.publicKey, data)
	②:originData,err :=   rsa.DecryptPKCS1v15(rand.Reader,key,cryptData)
   4、ras签名：私钥签名，公钥验签：
	①：signText,err := rsa.SignPKCS1v15(rand.Reader,key,crypto.hashMode,Hashed(data))
	②：err := rsa.VerifyPKCS1v15(&key.publicKey, crypto.hashMode, Hashed(data), signText )

*
*
*
 */
/*
RSA的Go语言API实现：
 私钥：type PrivateKey struct{
	PublicKey            // public part.
	D         *big.Int   // private exponent
	Primes    []*big.Int // prime factors of N, has >= 2 elements.

	// Precomputed contains precomputed values that speed up private
	// operations, if available.
	Precomputed PrecomputedValues
}

// A PublicKey represents the public part of an RSA key.
type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
}*/
//改函数用于生成一对秘钥
func CreatePairKeys() (*rsa.PrivateKey,  error) {
	//1、先生成私钥
	var bits int
	flag.IntVar(&bits, "b", 2048, "密钥用法")
	fmt.Println("bits的值：",bits)
	fmt.Printf("bits的类型：%T\n",bits)
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil,  err
	}
	//2、根据privateKey生成公钥
	//publicKey := privateKey.PublicKey
	// 3、 返回私钥和公钥
	return privateKey,  nil
}
func GenerateKeys(filename string) (error){
	pri,err := CreatePairKeys()
	if err != nil {
		return err
	}
	 return GeneratePemFilesByPrivateKey(pri,"wu")

}

/*+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++加密
RSA算法公钥对数据进行加密，返回加密的密文
*/
func RSAEncrypt(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

//RSA算法私钥对密文进行解密，返回解密后的数据
func RSADecrypt(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
}

/*++++++++++++++++++++++++++++++++++++++++++++++++++++++++++签名
RSA算法私钥对数据进行加密，返回加密的密文
iota
*/
func RSASign(privatKey *rsa.PrivateKey, data []byte) ([]byte, error) {

	return rsa.SignPKCS1v15(rand.Reader, privatKey, crypto.SHA256, HashSha256(data))
}

/***
*使用RSA算法对数据进行签名验证，并返回验证签名的结果
	验证通过，返回true，nil
	验证不通过，返回false，error
*/
func RSAVerify(publicKey *rsa.PublicKey, data, signText []byte) (bool, error) {
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256,HashSha256(data), signText, )
	return err == nil, err
}
func main() {
	data := "士大夫哈里发积分卡历史的角度考虑经济的"
	priv, err := CreatePairKeys()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pub := &priv.PublicKey
	//将私钥存入硬盘持久化存储
	err = GeneratePemFilesByPrivateKey(priv,"")
	if err != nil {
		fmt.Println("生成密钥失败：", err.Error())
	}
	cipherText, err := RSAEncrypt(pub, []byte(data))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(cipherText))
	originText, err := RSADecrypt(priv, cipherText)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(originText))
	fmt.Println("___________________________________________________签名")
	signText, err := RSASign(priv, []byte(data))
	if err != nil {
		fmt.Println("签名失败")
		fmt.Println(err.Error())
	}
	verifyResult, err := RSAVerify(pub, []byte(data), signText)
	if err != nil {
		fmt.Println("签名验证失败")
		fmt.Println(err.Error())
	}
	fmt.Println("验证结果为：", verifyResult)
}

/**
* 根据给定的私钥数据，生成私钥文件
 */
func GeneratePemFilesByPrivateKey(key *rsa.PrivateKey,filename string) error {
	//根据PCKS1规则，序列化的私钥
	priStream := x509.MarshalPKCS1PrivateKey(key)
	privateFile, err := os.Create("rsa_"+filename+"pri.pem") //new一个存私钥的文件
	if err != nil {
		return err
	}
	pubStream := x509.MarshalPKCS1PublicKey(&key.PublicKey)

	publicFile, err := os.Create("rsa_"+filename+"pub.pem")
	if err != nil {
		return err
	}
	block1 := &pem.Block{
		Type:    " RSA Public Key ",
		Bytes:   pubStream,
	}
	err = pem.Encode(publicFile, block1)
	if err != nil {
		return err
	}
	//pem文件中的格式 结构体   pem: 证书文件后缀
	block := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   priStream,
	}
	//调用pem包下的encode函数，将我们定义好的pem文件格式和内容写入到privateFile文件中
	err = pem.Encode(privateFile, block)
	if err != nil {
		return err
	}
	return nil
}

/**
*使用MD5哈希[]byte
 */
func HashMD5(data []byte) []byte {
	md5Hash := md5.New()
	md5Hash.Write(data)
	return md5Hash.Sum(nil)
}
func HashSha256(data []byte) []byte {
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
