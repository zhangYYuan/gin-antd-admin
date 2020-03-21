package dao

import (
	"github.com/jinzhu/gorm"
	"para/para-api/model"
)

type IProduct interface {
	Insert(product *model.Product)(int64, error)
	Delete(int64) bool
	Update(product *model.Product) error
	SelectById(int64)(*model.Product, error)
	
	SelectAll()([]*model.Product, error)
}

// 用这个结构体实现接口
type ProductManager struct {
	table string

}

// 构造函数实现代码自检

func NewProductManager(table string, db *gorm.DB) IProduct {
	return &ProductManager{}
}
