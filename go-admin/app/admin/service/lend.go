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
	var data2 models.Lend
	var render models.Render
	var book models.Book

	e.Orm.Where("book_id = ?", c.BookId).Where("render_id = ?", c.RenderId).Where("render_id = ?", c.RenderId).Where("state = ?", 0).First(&data2)
	e.Orm.Where("id = ?", c.RenderId).First(&render)
	e.Orm.Where("id = ?", c.BookId).First(&book)

	if data2.Id != 0 {
		err = errors.New("此书还未归还")
		e.Log.Errorf("LendService Insert error:%s \r\n", err)
		return err
	}

	if render.Count == 0 {
		err = errors.New("读者借阅数不足")
		e.Log.Errorf("LendService Insert error:%s \r\n", err)
		return err
	}

	if book.Count == 0 {
		err = errors.New("书籍库存不足了")
		e.Log.Errorf("LendService Insert error:%s \r\n", err)
		return err
	}

	render.Count = render.Count - 1
	book.Count = book.Count - 1

	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("LendService Insert error:%s \r\n", err)
		return err
	}
	e.Orm.Save(&render)
	e.Orm.Save(&book)
	return nil
}

// Update 修改Lend对象
func (e *Lend) Update(c *dto.LendUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Lend{}
	var render models.Render
	var book models.Book

	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())

	e.Orm.Where("id = ?", data.RenderId).First(&render)
	e.Orm.Where("id = ?", data.BookId).First(&book)

	c.Generate(&data)

	db := e.Orm.Save(&data)
	render.Count = render.Count + 1
	book.Count = book.Count + 1
	if err = db.Error; err != nil {
		e.Log.Errorf("LendService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	e.Orm.Save(&render)
	e.Orm.Save(&book)
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
