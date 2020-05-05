package service

import (
	"bee/db"
	"bee/model"
	"fmt"
)

func Resigster(u model.SysUser) (err error, userInter model.SysUser)  {
	//var user model.SysUser
	//notRegister := db.BeeDB.Not("username", u.UserName).First(&user)
	//fmt.Println(notRegister)
	fmt.Println(u.NickName)
	 er := db.BeeDB.Create(u)
	fmt.Println(er)
	return err, u
}