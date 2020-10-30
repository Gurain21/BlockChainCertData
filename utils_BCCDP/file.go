package utils_BCCDP

import (
	"io"
	"os"
)
//打开路径为filePath的文件，如果没有则自动创建该文件
func OpenFile(filePath string)(*os.File,error)  {
	file,err := os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		return nil,err
	}
	return file,nil
}
//保存文件，把file的内容保存到新创建的savefile文件中
func SaveFile(fileName string,file io.Reader)(int64,error)  {
	saveFile,err := OpenFile(fileName)
	if err != nil{
		return -1,err
	}
	length,err := io.Copy(saveFile,file)
	if err != nil {
		return -1,err
	}
	return length,nil
}
//打开目录，如果没有则创建目录
func OpenDir(dirPath string)(*os.File,error)  {
	dir,err := os.Open(dirPath)
	if err !=nil {
		dir ,err= os.Create(dirPath)
		if err != nil {
			return nil,err
		}
	}
	return dir,err
}