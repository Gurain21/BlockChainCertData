package main

import (
	"BlockChainCertDataPorject/utils_BCCDP"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	/**
	*ecc : elliptic curve cryptography
	*同等加密强度ecc的长度小于rsa
	*带宽低
	*ecdsa包实现 :elliptic curve digital signature algorithm
	 */
	priv,err := GenerateECDSAKey()
	if err != nil {
		fmt.Println("生成私钥失败：",err.Error())
		return
	}
	data := "花花画画画花花"
	r,s,err := ECDSASign(priv,[]byte(data))
	if err != nil {
		fmt.Println("ECDSA数字签名失败：",err.Error())
		return
	}
	if ECDSAVerify(priv.PublicKey,r,s,[]byte(data)){
		fmt.Println("ECDSA数字签名成功！")
	}else{
		fmt.Println("ECDSA数字签名失败！")
	}

}
//---------------------生成私钥和公钥的密钥对-------------------------

func GenerateECDSAKey()(*ecdsa.PrivateKey,error)  {
	return ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
}
//-----------------私钥签名，公钥验签----------------------
//______私钥签名__________________________________________-______________-
func ECDSASign(pri *ecdsa.PrivateKey,data []byte)(*big.Int,*big.Int,error){
	return  ecdsa.Sign(rand.Reader,pri,utils_BCCDP.SHA256HashByte(data))
}
//______公钥验签___________________________________________-_______________-____
func ECDSAVerify(pub ecdsa.PublicKey,r,s *big.Int,data []byte)(bool){
	return	ecdsa.Verify(&pub,utils_BCCDP.SHA256HashByte(data),r,s)
}
