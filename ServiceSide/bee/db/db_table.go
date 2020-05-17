package db

import "bee/model"

func InitTables() {
	BeeDB.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
		model.SysMenu{},
	)
}

