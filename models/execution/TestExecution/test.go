package TestExecution

import (
	"github.com/mdcaceres/doctest/models/execution"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TestExecution struct {
	gorm.Model
	ProjectID uint
	CaseID    uint
	Status    string
	UserId    uint
	Asserts   datatypes.JSONMap
}

func (e TestExecution) Execute() execution.IExecutable {
	return e
}
