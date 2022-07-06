package models

import (
	"blog-go/dao/mysql"
)

type User struct {
	UserID       string `json:"userid" gorm:"column:id"`
	UserName     string `json:"username" gorm:"column:name"`
	UserPassword string `json:"userpassword" gorm:"column:password"`
}

// 用户注册 Register

func Register(user User) (result string, err error) {
	res := mysql.DB.First(&user)
	if res.RowsAffected == 1 {
		result = "账号已存在"
	} else {
		err = mysql.DB.Create(&user).Error
		result = "注册信息可用"
	}

	return
}

// 用户登录 Login

func Login(user1 User) (result string, err error) {
	user := user1
	err = mysql.DB.First(&user).Error

	if user.UserPassword == user1.UserPassword {
		result = "登录成功"
	} else {
		result = "账号或密码错误"
	}

	return
}

// 补全用户信息 CUser
func CUser(user *User) (err error) {
	err = mysql.DB.First(&user).Error
	return err
}
