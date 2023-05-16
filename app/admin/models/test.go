package models

import (
	"go-admin/common/models"
)

type Test struct {
	models.Model

	Name string `json:"name" gorm:"type:varchar(128);comment:名称"`
	Dept string `json:"dept" gorm:"type:varchar(255);comment:部门"`
	models.ModelTime
	models.ControlBy
}

func (Test) TableName() string {
	return "test"
}

func (e *Test) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Test) GetId() interface{} {
	return e.Id
}
