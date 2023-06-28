package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RenderGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:render" comment:"读者用户名"`
	Gender         string `form:"gender"  search:"type:exact;column:gender;table:render" comment:"性别"`
	Phone          string `form:"phone"  search:"type:contains;column:phone;table:render" comment:"手机号"`
	RenderOrder
}

type RenderOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:render"`
	RenderId  string `form:"renderIdOrder"  search:"type:order;column:render_id;table:render"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:render"`
	Password  string `form:"passwordOrder"  search:"type:order;column:password;table:render"`
	Gender    string `form:"genderOrder"  search:"type:order;column:gender;table:render"`
	Phone     string `form:"phoneOrder"  search:"type:order;column:phone;table:render"`
	Count     int    `form:"countOrder"  search:"type:order;column:count;table:render"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:render"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:render"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:render"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:render"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:render"`
}

func (m *RenderGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RenderInsertReq struct {
	Id       int    `json:"-" comment:"读者id"` // 读者id
	RenderId string `json:"renderId" comment:"读者编号"`
	Name     string `json:"name" comment:"读者用户名"`
	Password string `json:"password" comment:"密码"`
	Gender   string `json:"gender" comment:"性别"`
	Phone    string `json:"phone" comment:"手机号"`
	Count    int    `json:"count" comment:"可借阅数量"`
	common.ControlBy
}

func (s *RenderInsertReq) Generate(model *models.Render) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RenderId = s.RenderId
	model.Name = s.Name
	model.Password = s.Password
	model.Gender = s.Gender
	model.Phone = s.Phone
	model.Count = s.Count
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *RenderInsertReq) GetId() interface{} {
	return s.Id
}

type RenderUpdateReq struct {
	Id       int    `uri:"id" comment:"读者id"` // 读者id
	RenderId string `json:"renderId" comment:"读者编号"`
	Name     string `json:"name" comment:"读者用户名"`
	Password string `json:"password" comment:"密码"`
	Gender   string `json:"gender" comment:"性别"`
	Phone    string `json:"phone" comment:"手机号"`
	Count    int    `json:"count" comment:"可借阅数量"`
	common.ControlBy
}

func (s *RenderUpdateReq) Generate(model *models.Render) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RenderId = s.RenderId
	model.Name = s.Name
	model.Password = s.Password
	model.Gender = s.Gender
	model.Phone = s.Phone
	model.Count = s.Count
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *RenderUpdateReq) GetId() interface{} {
	return s.Id
}

// RenderGetReq 功能获取请求参数
type RenderGetReq struct {
	Id int `uri:"id"`
}

func (s *RenderGetReq) GetId() interface{} {
	return s.Id
}

// RenderDeleteReq 功能删除请求参数
type RenderDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RenderDeleteReq) GetId() interface{} {
	return s.Ids
}
