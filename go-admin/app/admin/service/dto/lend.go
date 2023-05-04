package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type LendGetPageReq struct {
	dto.Pagination `search:"-"`
	RenderId       string `form:"renderId"  search:"type:exact;column:render_id;table:lend" comment:"读者id"`
	BookId         string `form:"bookId"  search:"type:contains;column:book_id;table:lend" comment:"图书id"`
	State          string `form:"state"  search:"type:exact;column:state;table:lend" comment:"状态"`
	LendOrder
}

type LendOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:lend"`
	RenderId  string `form:"renderIdOrder"  search:"type:order;column:render_id;table:lend"`
	BookId    string `form:"bookIdOrder"  search:"type:order;column:book_id;table:lend"`
	LendDate  string `form:"lendDateOrder"  search:"type:order;column:lend_date;table:lend"`
	BackDate  string `form:"backDateOrder"  search:"type:order;column:back_date;table:lend"`
	State     string `form:"stateOrder"  search:"type:order;column:state;table:lend"`
	Fine      string `form:"fineOrder"  search:"type:order;column:fine;table:lend"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:lend"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:lend"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:lend"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:lend"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:lend"`
}

func (m *LendGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type LendInsertReq struct {
	Id       int    `json:"-" comment:"借阅记录id"` // 借阅记录id
	RenderId string `json:"renderId" comment:"读者id"`
	BookId   string `json:"bookId" comment:"图书id"`
	LendDate string `json:"lendDate" comment:"借阅日期"`
	BackDate string `json:"backDate" comment:"最后归还日期"`
	State    string `json:"state" comment:"状态"`
	Fine     string `json:"fine" comment:"罚款"`
	common.ControlBy
}

func (s *LendInsertReq) Generate(model *models.Lend) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RenderId = s.RenderId
	model.BookId = s.BookId
	model.LendDate = s.LendDate
	model.BackDate = s.BackDate
	model.State = s.State
	model.Fine = s.Fine
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *LendInsertReq) GetId() interface{} {
	return s.Id
}

type LendUpdateReq struct {
	Id int `uri:"id" comment:"借阅记录id"` // 借阅记录id
	// RenderId string `json:"renderId" comment:"读者id"`
	// BookId   string `json:"bookId" comment:"图书id"`
	// LendDate string `json:"lendDate" comment:"借阅日期"`
	// BackDate string `json:"backDate" comment:"最后归还日期"`
	State string `json:"state" comment:"状态"`
	// Fine     string `json:"fine" comment:"罚款"`
	common.ControlBy
}

func (s *LendUpdateReq) Generate(model *models.Lend) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	// model.RenderId = s.RenderId
	// model.BookId = s.BookId
	// model.LendDate = s.LendDate
	// model.BackDate = s.BackDate
	model.State = s.State
	// model.Fine = s.Fine
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *LendUpdateReq) GetId() interface{} {
	return s.Id
}

// LendGetReq 功能获取请求参数
type LendGetReq struct {
	Id int `uri:"id"`
}

func (s *LendGetReq) GetId() interface{} {
	return s.Id
}

// LendDeleteReq 功能删除请求参数
type LendDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *LendDeleteReq) GetId() interface{} {
	return s.Ids
}
