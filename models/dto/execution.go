package dto

import (
	"github.com/mdcaceres/doctest/models/execution/TestExecution"
	"gorm.io/datatypes"
)

type TestExecutionRequest struct {
	ProjectID uint   `json:"project_id"`
	CaseID    uint   `json:"case_id"`
	Status    string `json:"status"`
	UserId    uint   `json:"user_id"`
	Asserts   datatypes.JSONMap
}

type TestExecutionResponse struct {
	ID          uint   `json:"id"`
	DateCreated string `json:"date_created"`
	ProjectID   uint   `json:"project_id"`
	CaseID      uint   `json:"case_id"`
	Status      string `json:"status"`
	UserId      uint   `json:"user_id"`
	Asserts     datatypes.JSONMap
}

func GetTestExecutionResponse(e *TestExecution.TestExecution) TestExecutionResponse {
	return TestExecutionResponse{
		ID:          e.ID,
		DateCreated: e.CreatedAt.String(),
		ProjectID:   e.ProjectID,
		CaseID:      e.CaseID,
		Status:      e.Status,
		UserId:      e.UserId,
		Asserts:     e.Asserts,
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

func GetSuiteExecutionResponse(e *TestExecution.TestExecution) TestExecutionResponse {
	return TestExecutionResponse{
		ID:          e.ID,
		DateCreated: e.CreatedAt.String(),
		ProjectID:   e.ProjectID,
		CaseID:      e.CaseID,
		Status:      e.Status,
		UserId:      e.UserId,
	}
}
