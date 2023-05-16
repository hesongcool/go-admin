package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TestGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:exact;column:name;table:test" comment:"名称"`
	Dept           string `form:"dept"  search:"type:exact;column:dept;table:test" comment:"部门"`
	TestOrder
}

type TestOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:test"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:test"`
	Dept      string `form:"deptOrder"  search:"type:order;column:dept;table:test"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:test"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:test"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:test"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:test"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:test"`
}

func (m *TestGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TestInsertReq struct {
	Id   int    `json:"-" comment:"主键编码"` // 主键编码
	Name string `json:"name" comment:"名称"`
	Dept string `json:"dept" comment:"部门"`
	common.ControlBy
}

func (s *TestInsertReq) Generate(model *models.Test) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Dept = s.Dept
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *TestInsertReq) GetId() interface{} {
	return s.Id
}

type TestUpdateReq struct {
	Id   int    `uri:"id" comment:"主键编码"` // 主键编码
	Name string `json:"name" comment:"名称"`
	Dept string `json:"dept" comment:"部门"`
	common.ControlBy
}

func (s *TestUpdateReq) Generate(model *models.Test) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Dept = s.Dept
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *TestUpdateReq) GetId() interface{} {
	return s.Id
}

// TestGetReq 功能获取请求参数
type TestGetReq struct {
	Id int `uri:"id"`
}

func (s *TestGetReq) GetId() interface{} {
	return s.Id
}

// TestDeleteReq 功能删除请求参数
type TestDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TestDeleteReq) GetId() interface{} {
	return s.Ids
}
