package utils_BCCDP

import (
	"errors"
	"fmt"
)


func CheckErrore (err error,str string)  {
	if err !=nil {
		fmt.Println(errors.New(str))
	}
}
func IsException(err error,detail string)  {
	defer func() {
		if ins, ok := recover().(error);ok{
			fmt.Println(detail,ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}