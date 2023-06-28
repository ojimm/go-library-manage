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

type Lend struct {
	api.Api
}

// GetPage 获取Lend列表
// @Summary 获取Lend列表
// @Description 获取Lend列表
// @Tags Lend
// @Param renderId query string false "读者id"
// @Param bookId query string false "图书id"
// @Param state query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Lend}} "{"code": 200, "data": [...]}"
// @Router /api/v1/lend [get]
// @Security Bearer
func (e Lend) GetPage(c *gin.Context) {
	req := dto.LendGetPageReq{}
	s := service.Lend{}
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
	list := make([]models.Lend, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Lend失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Lend
// @Summary 获取Lend
// @Description 获取Lend
// @Tags Lend
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Lend} "{"code": 200, "data": [...]}"
// @Router /api/v1/lend/{id} [get]
// @Security Bearer
func (e Lend) Get(c *gin.Context) {
	req := dto.LendGetReq{}
	s := service.Lend{}
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
	var object models.Lend

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Lend失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Lend
// @Summary 创建Lend
// @Description 创建Lend
// @Tags Lend
// @Accept application/json
// @Product application/json
// @Param data body dto.LendInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/lend [post]
// @Security Bearer
func (e Lend) Insert(c *gin.Context) {
	req := dto.LendInsertReq{}
	s := service.Lend{}
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
		e.Error(500, err, fmt.Sprintf("创建Lend失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Lend
// @Summary 修改Lend
// @Description 修改Lend
// @Tags Lend
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.LendUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/lend/{id} [put]
// @Security Bearer
func (e Lend) Update(c *gin.Context) {
	req := dto.LendUpdateReq{}
	s := service.Lend{}
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
		e.Error(500, err, fmt.Sprintf("归还Lend失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "归还成功")
}

// Delete 删除Lend
// @Summary 删除Lend
// @Description 删除Lend
// @Tags Lend
// @Param data body dto.LendDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/lend [delete]
// @Security Bearer
func (e Lend) Delete(c *gin.Context) {
	s := service.Lend{}
	req := dto.LendDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Lend失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
