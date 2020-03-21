package dao

import (
	"para.common/db"
	"para.common/models"
)

type UserDao struct {

}

func (user *UserDao) QueryByUsername (m models.User ){
	db.DB.First(&m)
}

func isTelephoneExist(phone string) bool {
	var user models.User
	db.DB.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}