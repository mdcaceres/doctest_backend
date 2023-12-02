package services

import (
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/models/execution/SuiteExecution"
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
	"github.com/mdcaceres/doctest/providers"
	"github.com/mdcaceres/doctest/services/mail"
)

type IExecutionService interface {
	SaveTestExecution(payload *dto.TestExecutionRequest) (*TestExecution.TestExecution, error)
	SaveSuiteExecution(suiteID uint) (*SuiteExecution.SuiteExecution, error)
}

type ExecutionService struct {
	ExecutionProvider providers.IExecutionProvider
	EmailService      *mail.EmailService
}

func (s ExecutionService) SaveTestExecution(payload *dto.TestExecutionRequest) (*TestExecution.TestExecution, error) {
	projectId := payload.ProjectID
	caseId := payload.CaseID
	userId := payload.UserID

	var steps []TestExecution.ExecutionStep

	for _, step := range payload.Steps {
		steps = append(steps, TestExecution.ExecutionStep{
			Order:       step.Order,
			Description: step.Description,
			Expected:    step.Expected,
			Status:      step.Status,
			Comment:     step.Comment,
		})
	}

	e := &TestExecution.TestExecution{
		ProjectID: projectId,
		CaseID:    caseId,
		UserId:    userId,
		Status:    payload.Status,
		Steps:     steps,
		Duration:  payload.Duration,
	}

	e, err := s.ExecutionProvider.Create(e)

	if err != nil {
		return nil, err
	}
	return e, nil
}

func (s *ExecutionService) SaveSuiteExecution(suiteID uint) (*SuiteExecution.SuiteExecution, error) {
	return nil, nil
}

func NewExecutionService() IExecutionService {
	return &ExecutionService{
		ExecutionProvider: providers.NewExecutionProvider(),
		EmailService:      mail.NewEmailService(),
	}
}
