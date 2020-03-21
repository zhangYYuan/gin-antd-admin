package dao

import (
	"para/para-api/model"
	"para.common/db"
)

// 添加
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	//
	db.DB.Create(&category)
	return category.CategoryId, nil
}

// 获取单个
func GetCategoryById(categoryId int64)(model.Category, error)  {
	var category = &model.Category{CategoryId:categoryId}

	db.DB.Find(&category)
	return *category, nil
}


// 获取多个

//func GetCategoryList(categoryIds []int64) (models []*model.Category, err error)  {
//	db.DB.Exec("select, id, category_name, category_no from category where id in (?)", categoryIds)
//	if err != nil {
//		return
//	}
//	return
//}

//func GetAllCategoryList() (models *model.Category, err error) {
//
//}