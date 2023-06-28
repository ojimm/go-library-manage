package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BookGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:book" comment:"图书名称"`
	Cid            string `form:"cid"  search:"type:exact;column:cid;table:book" comment:"书籍类别id"`
	BookOrder
}

type BookOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:book"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:book"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:book"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:book"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:book"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:book"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:book"`
	Author    string `form:"authorOrder"  search:"type:order;column:author;table:book"`
	Publish   string `form:"publishOrder"  search:"type:order;column:publish;table:book"`
	Isbn      string `form:"isbnOrder"  search:"type:order;column:isbn;table:book"`
	Intro     string `form:"introOrder"  search:"type:order;column:intro;table:book"`
	Language  string `form:"languageOrder"  search:"type:order;column:language;table:book"`
	Price     string `form:"priceOrder"  search:"type:order;column:price;table:book"`
	Pubdate   string `form:"pubdateOrder"  search:"type:order;column:pubdate;table:book"`
	Count     string `form:"countOrder"  search:"type:order;column:count;table:book"`
	Cid       string `form:"cidOrder"  search:"type:order;column:cid;table:book"`
}

func (m *BookGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BookInsertReq struct {
	Id       int    `json:"-" comment:"图书id"` // 图书id
	Name     string `json:"name" comment:"图书名称"`
	Author   string `json:"author" comment:"作者"`
	Publish  string `json:"publish" comment:"出版社"`
	Isbn     string `json:"isbn" comment:"ISBN"`
	Intro    string `json:"intro" comment:"简介"`
	Language string `json:"language" comment:"语言"`
	Price    string `json:"price" comment:"价格"`
	Pubdate  string `json:"pubdate" comment:"出版日期"`
	Count    int    `json:"count" comment:"库存"`
	Cid      string `json:"cid" comment:"书籍类别id"`
	common.ControlBy
}

func (s *BookInsertReq) Generate(model *models.Book) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Author = s.Author
	model.Publish = s.Publish
	model.Isbn = s.Isbn
	model.Intro = s.Intro
	model.Language = s.Language
	model.Price = s.Price
	model.Pubdate = s.Pubdate
	model.Count = s.Count
	model.Cid = s.Cid
}

func (s *BookInsertReq) GetId() interface{} {
	return s.Id
}

type BookUpdateReq struct {
	Id       int    `uri:"id" comment:"图书id"` // 图书id
	Name     string `json:"name" comment:"图书名称"`
	Author   string `json:"author" comment:"作者"`
	Publish  string `json:"publish" comment:"出版社"`
	Isbn     string `json:"isbn" comment:"ISBN"`
	Intro    string `json:"intro" comment:"简介"`
	Language string `json:"language" comment:"语言"`
	Price    string `json:"price" comment:"价格"`
	Pubdate  string `json:"pubdate" comment:"出版日期"`
	Count    int    `json:"count" comment:"库存"`
	Cid      string `json:"cid" comment:"书籍类别id"`
	common.ControlBy
}

func (s *BookUpdateReq) Generate(model *models.Book) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Author = s.Author
	model.Publish = s.Publish
	model.Isbn = s.Isbn
	model.Intro = s.Intro
	model.Language = s.Language
	model.Price = s.Price
	model.Pubdate = s.Pubdate
	model.Count = s.Count
	model.Cid = s.Cid
}

func (s *BookUpdateReq) GetId() interface{} {
	return s.Id
}

// BookGetReq 功能获取请求参数
type BookGetReq struct {
	Id int `uri:"id"`
}

func (s *BookGetReq) GetId() interface{} {
	return s.Id
}

// BookDeleteReq 功能删除请求参数
type BookDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BookDeleteReq) GetId() interface{} {
	return s.Ids
}
