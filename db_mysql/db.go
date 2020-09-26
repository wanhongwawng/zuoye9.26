package db_mysql

import (
	"beego02/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据连接错误，请检查配置")
	}
	//为全局变量赋值
	DB = db
}

//将用户信息保存到数据库表当中
func InserUser(user models.User) (int64, error) {
	//1,将用户密码进行hash脱敏，使用md5计算密码hash并存储hash值
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名：",user.Nick,
	"密码：",user.Password)
	result, err := DB.Exec("insert into user(nick,password)values(?,?)", user.Nick,user.Password)
	if err !=nil{//保存数据时遇到错误
		return -1,err
	}
	id,err :=result.RowsAffected()
	if  err !=nil {
		return -1,err
	}
	return  id,nil
}
//查询用户
func QueryUser()  {
	DB.QueryRow("select *from ")
}
