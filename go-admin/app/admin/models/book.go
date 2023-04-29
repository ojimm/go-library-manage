package models

import (
	"go-admin/common/models"
)

type Book struct {
	models.Model

	Name     string `json:"name" gorm:"type:varchar(128);comment:图书名称"`
	Author   string `json:"author" gorm:"type:varchar(255);comment:作者"`
	Publish  string `json:"publish" gorm:"type:varchar(255);comment:出版社"`
	Isbn     string `json:"isbn" gorm:"type:varchar(255);comment:ISBN"`
	Intro    string `json:"intro" gorm:"type:varchar(255);comment:简介"`
	Language string `json:"language" gorm:"type:varchar(255);comment:语言"`
	Price    string `json:"price" gorm:"type:double;comment:价格"`
	Pubdate  string `json:"pubdate" gorm:"type:varchar(255);comment:出版日期"`
	Count    string `json:"count" gorm:"type:int(11);comment:库存"`
	Cid      string `json:"cid" gorm:"type:int(11);comment:书籍类别id"`
	models.ModelTime
	models.ControlBy
}

func (Book) TableName() string {
	return "book"
}

func (e *Book) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Book) GetId() interface{} {
	return e.Id
}
