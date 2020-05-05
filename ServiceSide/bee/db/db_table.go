package db

import "bee/model"

func InitTables() {
	BeeDB.AutoMigrate(
		model.SysUser{},
	)
}