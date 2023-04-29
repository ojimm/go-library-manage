package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BookClassGetPageReq struct {
	dto.Pagination `search:"-"`
	BookClassOrder
}

type BookClassOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:book_class"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:book_class"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:book_class"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:book_class"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:book_class"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:book_class"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:book_class"`
}

func (m *BookClassGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BookClassInsertReq struct {
	Id   int    `json:"-" comment:"书籍类别id"` // 书籍类别id
	Name string `json:"name" comment:"书籍类别名称"`
	common.ControlBy
}

func (s *BookClassInsertReq) Generate(model *models.BookClass) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BookClassInsertReq) GetId() interface{} {
	return s.Id
}

type BookClassUpdateReq struct {
	Id   int    `uri:"id" comment:"书籍类别id"` // 书籍类别id
	Name string `json:"name" comment:"书籍类别名称"`
	common.ControlBy
}

func (s *BookClassUpdateReq) Generate(model *models.BookClass) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BookClassUpdateReq) GetId() interface{} {
	return s.Id
}

// BookClassGetReq 功能获取请求参数
type BookClassGetReq struct {
	Id int `uri:"id"`
}

func (s *BookClassGetReq) GetId() interface{} {
	return s.Id
}

// BookClassDeleteReq 功能删除请求参数
type BookClassDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BookClassDeleteReq) GetId() interface{} {
	return s.Ids
}
