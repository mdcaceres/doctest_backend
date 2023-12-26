package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
	"gorm.io/gorm"
	"time"
)

type IExecutionProvider interface {
	Create(t *TestExecution.TestExecution) (*TestExecution.TestExecution, error)
	GetByProjectIdStatusAndDateRange(projectId string, status string, start time.Time, end time.Time) ([]TestExecution.TestExecution, error)
	GetCountByProjectId(projectId string, start time.Time, end time.Time) (int64, error)
}

type ExecutionProvider struct {
	DB *gorm.DB
}

func NewExecutionProvider() IExecutionProvider {
	return &ExecutionProvider{
		DB: datasource.GetDB(),
	}
}

func (e *ExecutionProvider) Create(t *TestExecution.TestExecution) (*TestExecution.TestExecution, error) {
	result := e.DB.Create(t)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error creating new TestExecution case in our database [error:%v]", result.Error))
	}
	return t, nil
}

func (e *ExecutionProvider) GetByProjectIdStatusAndDateRange(projectId string, status string, start time.Time, end time.Time) ([]TestExecution.TestExecution, error) {
	var executions []TestExecution.TestExecution
	result := e.DB.Where("project_id = ? AND status = ? AND created_at BETWEEN ? AND ?", projectId, status, start, end).Find(&executions)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting TestExecution in our database [error:%v]", result.Error))
	}
	return executions, nil
}

func (e *ExecutionProvider) GetCountByProjectId(projectId string, start time.Time, end time.Time) (int64, error) {
	var count int64
	result := e.DB.Model(TestExecution.TestExecution{}).Where("project_id = ? AND created_at BETWEEN ? AND ?", projectId, start.Local(), end.Local()).Count(&count)

	if result.Error != nil {
		return 0, errors.New(fmt.Sprintf("error getting TestExecution in our database [error:%v]", result.Error))
	}
	return count, nil
}
