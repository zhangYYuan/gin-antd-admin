package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

/**
 权限
 */
type SysAuthority struct {
	CreatedAt time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId 	string			`json:"authorityId" gorm:"not null;unique;primary_key"`
	AuthorityName   string         `json:"authorityName"`
	ParentId        string         `json:"parentId"`
	//DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;association_jointable_foreignkey:data_authority_id"`
	//Children        []SysAuthority `json:"children"`
	//SysBaseMenus    []SysAuthority  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}

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
	Authority SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	AuthorityId string		`json:"authorityId" gorm:"default: 888"`
}
