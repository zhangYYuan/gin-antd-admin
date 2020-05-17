package model

import (
	"github.com/jinzhu/gorm"
)

type SysMenu struct {
	gorm.Model
	MenuLevel 	uint            `json:"menuLevel"`
	ParentId	string          `json:"parentId"`
	Path        string          `json:"path"`
	Name 		string          `json:"name"`
	Icon		string			`json:"icon"`
	Hidden 		bool          	`json:"hidden"`
	Component	string          `json:"component"`
	Sort 		int             `json:"sort"`
	SysAuthority []SysAuthority `json:"authoritis" gorm:"many2many:sys_authority_menus;"`
	Children 	 []SysMenu      `json:"children"`
}