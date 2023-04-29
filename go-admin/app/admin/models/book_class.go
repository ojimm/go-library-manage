package models

import (
	"go-admin/common/models"
)

type BookClass struct {
	models.Model

	Name string `json:"name" gorm:"type:varchar(128);comment:书籍类别名称"`
	models.ModelTime
	models.ControlBy
}

func (BookClass) TableName() string {
	return "book_class"
}

func (e *BookClass) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BookClass) GetId() interface{} {
	return e.Id
}
