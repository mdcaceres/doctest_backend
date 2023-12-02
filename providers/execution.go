package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
	"gorm.io/gorm"
)

type IExecutionProvider interface {
	Create(t *TestExecution.TestExecution) (*TestExecution.TestExecution, error)
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
