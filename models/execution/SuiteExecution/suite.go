package SuiteExecution

import (
	"github.com/mdcaceres/doctest/models/execution"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SuiteExecution struct {
	gorm.Model
	ProjectID  uint
	Status     string
	UserId     uint
	TestPassed datatypes.JSONMap
}

func (s SuiteExecution) Execute() execution.IExecutable {
	return s
}
