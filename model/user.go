package model

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"online_shop/service/jwtx"
)

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
}

func (u *User) Register(mobile, password string) (string, error) {
	var uid int
	var user User
	if !DB.Where("mobile = ? ", mobile).First(&user).RecordNotFound() {
		return "", errors.New("该手机已经注册")
	}
	newUser := User{
		Mobile:   mobile,
		Password: password,
	}
	if err := DB.Create(&newUser).Error; err != nil {
		return "", err
	}
	uid = newUser.Id
	tokenString, err := newUser.GenToken(uid)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *User) Login(mobile, password string) (string, error) {
	user := User{
		Mobile:   mobile,
		Password: password,
	}
	if DB.First(&user, &user).RecordNotFound() {
		return "", errors.New("请检查手机号和密码")
	}
	tokenString, err := user.GenToken(user.Id)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func (u *User) GenToken(id int) (string, error) {
	jwtParams := make(map[string]interface{})
	jwtParams["uid"] = id
	tokenString, err := jwtx.GenToken(jwtParams)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return tokenString, nil
}

func (u *User) Info(uid int) (*User, error) {
	var user User
	if DB.First(&user, uid).RecordNotFound() {
		return nil, errors.New("uid " + string(uid) + "NotFound")
	}
	return &user, nil
}
