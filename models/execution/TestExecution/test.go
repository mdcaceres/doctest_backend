package TestExecution

import (
	"github.com/mdcaceres/doctest/models/execution"
	"gorm.io/gorm"
	"time"
)

type TestExecution struct {
	gorm.Model
	ProjectID uint
	CaseID    uint
	Status    string
	UserId    uint
	Duration  time.Duration
	Steps     []ExecutionStep
}

func (e TestExecution) Execute() execution.IExecutable {
	return e
}
