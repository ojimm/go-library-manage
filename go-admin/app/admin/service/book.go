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

type Book struct {
	service.Service
}

// GetPage 获取Book列表
func (e *Book) GetPage(c *dto.BookGetPageReq, p *actions.DataPermission, list *[]models.Book, count *int64) error {
	var err error
	var data models.Book

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BookService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Book对象
func (e *Book) Get(d *dto.BookGetReq, p *actions.DataPermission, model *models.Book) error {
	var data models.Book

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBook error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Book对象
func (e *Book) Insert(c *dto.BookInsertReq) error {
    var err error
    var data models.Book
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BookService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Book对象
func (e *Book) Update(c *dto.BookUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Book{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("BookService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Book
func (e *Book) Remove(d *dto.BookDeleteReq, p *actions.DataPermission) error {
	var data models.Book

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveBook error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
