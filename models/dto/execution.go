package dto

import (
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
	"time"
)

type TestExecutionRequest struct {
	ProjectID uint                   `json:"project_id"`
	CaseID    uint                   `json:"case_id"`
	Status    string                 `json:"status"`
	UserID    uint                   `json:"user_id"`
	Steps     []StepExecutionRequest `json:"steps"`
	Duration  time.Duration          `json:"duration"`
}

type StepExecutionRequest struct {
	Order       uint   `json:"number"`
	Description string `json:"description"`
	Expected    string `json:"expected"`
	Status      string `json:"status"`
	Comment     string `json:"comment"`
}

type TestExecutionResponse struct {
	ID          uint   `json:"id"`
	DateCreated string `json:"date_created"`
	ProjectID   uint   `json:"project_id"`
	CaseID      uint   `json:"case_id"`
	Status      string `json:"status"`
	UserId      uint   `json:"user_id"`
	Steps       []StepExecutionResponse
}

type StepExecutionResponse struct {
	ID          uint   `json:"id"`
	Order       uint   `json:"order"`
	Description string `json:"description"`
	Result      string `json:"result"`
	CaseID      uint
	Status      string `json:"status"`
	Comment     string `json:"comment"`
}

func GetTestExecutionResponse(t *TestExecution.TestExecution) TestExecutionResponse {
	var s []StepExecutionResponse

	for _, step := range t.Steps {
		s = append(s, GetStepExecutionResponse(&step))
	}

	return TestExecutionResponse{
		ID:          t.ID,
		DateCreated: t.CreatedAt.String(),
		ProjectID:   t.ProjectID,
		CaseID:      t.CaseID,
		Status:      t.Status,
		UserId:      t.UserId,
		Steps:       s,
	}
}

func GetStepExecutionResponse(s *TestExecution.ExecutionStep) StepExecutionResponse {
	return StepExecutionResponse{
		ID:          s.ID,
		Order:       s.Order,
		Description: s.Description,
		Result:      s.Expected,
		CaseID:      s.CaseID,
		Status:      s.Status,
		Comment:     s.Comment,
	}
}

type SuiteExecutionRequest struct {
	ProjectID uint   `json:"project_id"`
	SuiteID   uint   `json:"suite_id"`
	Status    string `json:"status"`
	UserId    uint   `json:"user_id"`
}

type SuiteExecutionResponse struct {
	ID          uint   `json:"id"`
	DateCreated string `json:"date_created"`
	ProjectID   uint   `json:"project_id"`
	CaseID      uint   `json:"case_id"`
	Status      string `json:"status"`
	UserId      uint   `json:"user_id"`
}
