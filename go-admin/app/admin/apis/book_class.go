package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type BookClass struct {
	api.Api
}

// GetPage 获取BookClass列表
// @Summary 获取BookClass列表
// @Description 获取BookClass列表
// @Tags BookClass
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.BookClass}} "{"code": 200, "data": [...]}"
// @Router /api/v1/book-class [get]
// @Security Bearer
func (e BookClass) GetPage(c *gin.Context) {
    req := dto.BookClassGetPageReq{}
    s := service.BookClass{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.BookClass, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取BookClass失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取BookClass
// @Summary 获取BookClass
// @Description 获取BookClass
// @Tags BookClass
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.BookClass} "{"code": 200, "data": [...]}"
// @Router /api/v1/book-class/{id} [get]
// @Security Bearer
func (e BookClass) Get(c *gin.Context) {
	req := dto.BookClassGetReq{}
	s := service.BookClass{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.BookClass

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取BookClass失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建BookClass
// @Summary 创建BookClass
// @Description 创建BookClass
// @Tags BookClass
// @Accept application/json
// @Product application/json
// @Param data body dto.BookClassInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/book-class [post]
// @Security Bearer
func (e BookClass) Insert(c *gin.Context) {
    req := dto.BookClassInsertReq{}
    s := service.BookClass{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建BookClass失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改BookClass
// @Summary 修改BookClass
// @Description 修改BookClass
// @Tags BookClass
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BookClassUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/book-class/{id} [put]
// @Security Bearer
func (e BookClass) Update(c *gin.Context) {
    req := dto.BookClassUpdateReq{}
    s := service.BookClass{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改BookClass失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除BookClass
// @Summary 删除BookClass
// @Description 删除BookClass
// @Tags BookClass
// @Param data body dto.BookClassDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/book-class [delete]
// @Security Bearer
func (e BookClass) Delete(c *gin.Context) {
    s := service.BookClass{}
    req := dto.BookClassDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除BookClass失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
