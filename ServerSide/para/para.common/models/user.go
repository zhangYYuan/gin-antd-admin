package models

import (
	"github.com/jinzhu/gorm"
)

type SysUser struct {
	gorm.Model
	UserName string `json:"userName"`
	PassWord string	`json:"passWord"`
}

//func (u *SysUser) Login(err error, user *SysUser) {
//	var user SysUser
//	u.PassWord = utils.MD5V([]byte(u.PassWord))
//	return use
//}