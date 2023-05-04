package models

import (
	"go-admin/common/models"
)

type Lend struct {
	models.Model

	RenderId string `json:"renderId" gorm:"type:int(11);comment:读者id"`
	BookId   string `json:"bookId" gorm:"type:int(11);comment:图书id"`
	LendDate string `json:"lendDate" gorm:"type:varchar(255);comment:借阅日期"`
	BackDate string `json:"backDate" gorm:"type:varchar(255);comment:最后归还日期"`
	State    string `json:"state" gorm:"type:int(11);comment:状态"`
	Fine     string `json:"fine" gorm:"type:double;comment:罚款"`
	models.ModelTime
	models.ControlBy
}

func (Lend) TableName() string {
	return "lend"
}

func (e *Lend) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Lend) GetId() interface{} {
	return e.Id
}
