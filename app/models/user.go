package models


import (
	"demo_gorm/configs/db"
	"fmt"
)

type User struct{
	Id int `json:"id"`
	Email string
}

type userModel struct{}

var UserModel userModel

func init() {
	UserModel = userModel{}	
}

func AuthenticatePassword(password string)string{
	var user User
	db.MysqlConn.Where(&User{Email: "pratap.raudra@gmail.com"}).First(&user)
	return user.Email
}

func (um *userModel)FetchByEmail(email string) User{
	var user User
	db.MysqlConn.Where(&User{Email: email}).First(&user)	
	fmt.Printf("%s", user)
	return user
}