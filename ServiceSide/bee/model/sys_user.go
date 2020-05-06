package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

/**
 用户模型
 */
type SysUser struct {
	gorm.Model
	UUID uuid.UUID 		`json:"uuid"`
	UserName string		`json:"userName"`
	Password string		`json:"-"`
	NickName string		`json:"nickName" gorm:"default:'cute bee'"`
	HeaderImg string    `json:"headerImg" gorm:"default:'https://jdc.jd.com/img/200'"`

}
