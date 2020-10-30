//该go文件用于定义转化类函数
package utils_BCCDP

import (
	"bytes"
	"encoding/binary"
)
//将int64类型转换未[]byte类型
func Int64ToByte(num int64)([]byte,error)  {
	buff := new(bytes.Buffer)
	err :=binary.Write(buff,binary.LittleEndian,num)
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}

