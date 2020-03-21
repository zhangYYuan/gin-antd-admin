package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Id int64
	CategoryId int64
	Title string
	Summary string
	ViewCount uint
	CommentCount uint
	UserName string
}



type ArticleDetail struct {
	Article
	Content string
}

type ArticleRecord struct {
	Article
	Category
}