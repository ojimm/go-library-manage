package models

import (
	"go-admin/common/models"
)

type Render struct {
	models.Model

	RenderId string `json:"renderId" gorm:"type:int(11);comment:读者编号"`
	Name     string `json:"name" gorm:"type:varchar(255);comment:读者用户名"`
	Password string `json:"password" gorm:"type:varchar(255);comment:密码"`
	Gender   string `json:"gender" gorm:"type:varchar(255);comment:性别"`
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:手机号"`
	Count    int    `json:"count" gorm:"type:int(11);comment:可借阅数量"`
	models.ModelTime
	models.ControlBy
}

func (Render) TableName() string {
	return "render"
}

func (e *Render) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Render) GetId() interface{} {
	return e.Id
}
