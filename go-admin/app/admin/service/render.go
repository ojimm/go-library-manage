package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Render struct {
	service.Service
}

// GetPage 获取Render列表
func (e *Render) GetPage(c *dto.RenderGetPageReq, p *actions.DataPermission, list *[]models.Render, count *int64) error {
	var err error
	var data models.Render

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RenderService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Render对象
func (e *Render) Get(d *dto.RenderGetReq, p *actions.DataPermission, model *models.Render) error {
	var data models.Render

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRender error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Render对象
func (e *Render) Insert(c *dto.RenderInsertReq) error {
    var err error
    var data models.Render
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RenderService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Render对象
func (e *Render) Update(c *dto.RenderUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Render{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("RenderService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Render
func (e *Render) Remove(d *dto.RenderDeleteReq, p *actions.DataPermission) error {
	var data models.Render

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveRender error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
