package services

import (
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/models/execution/SuiteExecution"
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
)

type IExecutionService interface {
	ExecuteCase(payload dto.TestExecutionRequest) (*TestExecution.TestExecution, error)
	ExecuteSuite(suiteID uint) (*SuiteExecution.SuiteExecution, error)
}
type ExecutionService struct {
}

func (s *ExecutionService) ExecuteCase(caseID uint) (*TestExecution.TestExecution, error) {
	return nil, nil
}

func (s *ExecutionService) ExecuteSuite(suiteID uint) (*SuiteExecution.SuiteExecution, error) {
	return nil, nil
}
