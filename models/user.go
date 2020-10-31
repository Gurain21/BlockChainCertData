package models

import (
	"BlockChainCertDataPorject/database_mysql"
	"BlockChainCertDataPorject/utils_BCCDP"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"passworld"`
	Name string `form:"name"`
	Card string `form:"card"`
	Sex string `form:"sex"`
}
//添加用户
func(u User)AddUser()(int64,error)  {
	u.Password = utils_BCCDP.MD5HashString(u.Password)
	result,err := database_mysql.DB_BCCDP.Exec("insert into user(phone,password)values (?,?)",u.Phone,u.Password)
	if err != nil {
		return  -1,err
	}
	rows,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return rows,nil
}
//通过phone和password 来判断登录的用户是否存在我们的本地数据库
func (u User)QueryUser()(*User,error)  {
	u.Password = utils_BCCDP.MD5HashString(u.Password)
	row := database_mysql.DB_BCCDP.QueryRow("select id,phone,password,name,sex,card from user where phone =?and password = ?",u.Phone,u.Passworld)
	err := row.Scan(&u.Id,&u.Phone,&u.Password,&u.Name,&u.Sex,&u.Card)
	if err != nil {
		return nil,err
	}
	return &u,nil
}

func (u User) QueryUserIdByPhone() (*User, error) {
	row := database_mysql.DB_BCCDP.QueryRow("select id from user where phone = ?", u.Phone)
	var user User
	err := row.Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}















//用户注册信息完善
func (u User)UpdataUser()(int64,error)  {
	result,err:=database_mysql.DB_BCCDP.Exec("updata user set name = ?,card = ?,sex =? where phone = ?",u.Name,u.Card,u.Sex,u.Phone)
	if err != nil {
		return -1,err
	}
	rows,err := result.RowsAffected()//RowsAffected 返回影响的行数和error
	if err != nil {
		return -1,err
	}
	return rows,nil
}
