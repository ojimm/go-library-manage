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

type Lend struct {
	service.Service
}

// GetPage 获取Lend列表
func (e *Lend) GetPage(c *dto.LendGetPageReq, p *actions.DataPermission, list *[]models.Lend, count *int64) error {
	var err error
	var data models.Lend

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("LendService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Lend对象
func (e *Lend) Get(d *dto.LendGetReq, p *actions.DataPermission, model *models.Lend) error {
	var data models.Lend

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetLend error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Lend对象
func (e *Lend) Insert(c *dto.LendInsertReq) error {
	var err error
	var data models.Lend
	c.Fine = "0"
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("LendService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Lend对象
func (e *Lend) Update(c *dto.LendUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Lend{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("LendService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Lend
func (e *Lend) Remove(d *dto.LendDeleteReq, p *actions.DataPermission) error {
	var data models.Lend

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveLend error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
