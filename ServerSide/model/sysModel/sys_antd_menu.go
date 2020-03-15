package sysModel

import (
	"errors"
	"fmt"
	"gin-vue-admin/init/qmsql"
	"github.com/jinzhu/gorm"
)

type SysAntdMenu struct {
	gorm.Model
	MenuLevel uint          `json:"_"`
	ParentId  string        `json:"parentId"`
	Path      string        `json:"path"`
	Component string		`json:"component"`
	Name      string        `json:"name"`
	Icon      string        `json:"icon"`
	Routes    []SysBaseMenu `json:"routes"`
}

func (m *SysAntdMenu) AddAntdMenu() (err error) {
	fmt.Println(m.Name)
	findOne := qmsql.DEFAULTDB.Where("name = ?", m.Name).Find(&SysBaseMenu{}).Error
	if findOne != nil {
		err = qmsql.DEFAULTDB.Create(m).Error
	} else {
		err = errors.New("name已存")
	}
	return err
}