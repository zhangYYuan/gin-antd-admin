package service

import (
	"bee/db"
	"bee/model"
	"errors"
)

func AddMenu(menu model.SysMenu) (err error) {
	findOne := db.BeeDB.Where("name = ?", menu.Name).Find(&model.SysMenu{}).Error
	if findOne != nil {
		 err = db.BeeDB.Create(&menu).Error
	} else {
		err = errors.New("菜单名已存在")
	}
	return err
}