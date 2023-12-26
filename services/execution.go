package services

import (
	"errors"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/models/execution/SuiteExecution"
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
	"github.com/mdcaceres/doctest/providers"
	"github.com/mdcaceres/doctest/services/mail"
	"time"
)

type IExecutionService interface {
	SaveTestExecution(payload *dto.TestExecutionRequest) (*TestExecution.TestExecution, error)
	SaveSuiteExecution(suiteID uint) (*SuiteExecution.SuiteExecution, error)
	GetPercentage(projectId string, status string, start time.Time, end time.Time) (float64, error)
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

func (s *ExecutionService) GetPercentage(projectId string, status string, start time.Time, end time.Time) (float64, error) {
	count, err := s.ExecutionProvider.GetCountByProjectId(projectId, start, end)
	if err != nil {
		return 0, err
	}
	filtered, err := s.ExecutionProvider.GetByProjectIdStatusAndDateRange(projectId, status, start, end)
	if err != nil {
		return 0, err
	}

	ff := float64(len(filtered))
	fc := float64(count)

	if fc == 0 {
		return 0, errors.New("cant divide by 0")
	}
	d := ff / fc
	return d * float64(100), nil
}
