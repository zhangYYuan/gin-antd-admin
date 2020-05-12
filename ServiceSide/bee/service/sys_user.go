package service

import (
	"bee/db"
	"bee/model"
	"bee/utils"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func Register(u model.SysUser) (error, model.SysUser)  {
	var user model.SysUser
	var err error
	notRegister := db.BeeDB.Where("user_name = ?", u.UserName).First(&user).RecordNotFound()
	fmt.Println(notRegister)
	if !notRegister {
		return errors.New("用户名已注册"), user
	} else {
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = db.BeeDB.Create(&u).Error
	}
	return err, u
}

func Login(u *model.SysUser) (error, *model.SysUser)  {
	var user model.SysUser
	var err error
	u.Password = utils.MD5V([]byte(u.Password))
	err = db.BeeDB.Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user).Error
	if err != nil {
		return err, &user
	}
	err = db.BeeDB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}