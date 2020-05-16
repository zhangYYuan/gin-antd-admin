package service

import (
	"bee/db"
	"bee/global"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"strings"
)


func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(db.BeeDB)
	e := casbin.NewEnforcer(global.BeeConfig.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error)  {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return (bool)(ParamsMatch(name1, name2)), nil
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}