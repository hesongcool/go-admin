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

type Test struct {
	service.Service
}

// GetPage 获取Test列表
func (e *Test) GetPage(c *dto.TestGetPageReq, p *actions.DataPermission, list *[]models.Test, count *int64) error {
	var err error
	var data models.Test

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TestService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Test对象
func (e *Test) Get(d *dto.TestGetReq, p *actions.DataPermission, model *models.Test) error {
	var data models.Test

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTest error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Test对象
func (e *Test) Insert(c *dto.TestInsertReq) error {
    var err error
    var data models.Test
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TestService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Test对象
func (e *Test) Update(c *dto.TestUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Test{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("TestService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Test
func (e *Test) Remove(d *dto.TestDeleteReq, p *actions.DataPermission) error {
	var data models.Test

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTest error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
