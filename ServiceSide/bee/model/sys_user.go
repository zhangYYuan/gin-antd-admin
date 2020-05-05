package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID uuid.UUID `json:"uuid"`
	UserName string	`json:"userName"`
	Password string	`json:"password"`
	NickName string	`json:"nickName"`
	HeaderImg string	`json:"headerImg"`

}
