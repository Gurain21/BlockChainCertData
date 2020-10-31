package models

import (
	"BlockChainCertDataPorject/database_mysql"
	"BlockChainCertDataPorject/utils_BCCDP"
)

type UploadRecord struct {
	Id             int
	UserId         int
	FileName       string
	FileSize       int64
	FileCert       string //认证号:文件的 md5哈希值
	FileTitle      string
	CertTime       int64
	CertTimeFormat string //仅作为格式化展示使用的字段
}

//保存认证文件的数据到数据库 表中
func (u UploadRecord) SaveRecord() (int64, error) {
	result, err := database_mysql.DB_BCCDP.Exec("insert into upload_record(user_id, file_name, file_size, file_cert, file_title, cert_time) "+
		"values(?,?,?,?,?,?) ", u.UserId, u.FileName, u.FileSize, u.FileCert, u.FileTitle, u.CertTime)
	if err != nil {
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rows, nil
}

//通过用户id来查询用户存证的文件数据
func QueryRecordsByUserId(userId int) ([]UploadRecord, error) {
	result, err := database_mysql.DB_BCCDP.Query("select id, user_id, file_name, file_size, file_cert, file_title, cert_time from upload_record where user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	//从rs中读取查询到的数据，返回
	records := make([]UploadRecord, 0) //容器
	for result.Next() {
		var record UploadRecord
		err := result.Scan(&record.Id, &record.UserId, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime)
		if err != nil {
			return nil, err
		}
		//整形 --> 字符串:xxxx年mm月dd日 hh:MM:ss
		tStr := utils_BCCDP.TimeFormat(record.CertTime, utils_BCCDP.TIME_FORMAT_THREE)
		record.CertTimeFormat = tStr
		records = append(records, record)
	}
	return records, nil
}
