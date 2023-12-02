package dto

import (
	"github.com/mdcaceres/doctest/models"
	"time"
)

type CaseRequest struct {
	CreatorId   string              `json:"creator_id"`
	Title       string              `json:"title"`
	Type        string              `json:"type"`
	SuiteId     uint                `json:"suite_id"`
	Priority    string              `json:"priority"`
	Description string              `json:"description"`
	Steps       []map[string]string `json:"steps"`
	Duration    string              `json:"duration"`
	ProjectId   string              `json:"project_id"`
	Status      string              `json:"status"`
}

type CaseResponse struct {
	ID          uint          `json:"id"`
	CreatorId   uint          `json:"creator_id"`
	Title       string        `json:"title"`
	Type        string        `json:"type"`
	SuiteId     uint          `json:"suite_id"`
	Priority    string        `json:"priority"`
	Description string        `json:"description"`
	Steps       []models.Step `json:"steps"`
	Duration    time.Duration `json:"duration"`
	Status      string        `json:"status"`
	ProjectId   uint          `json:"project_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func GetCaseResponse(testCase *models.Case) CaseResponse {
	return CaseResponse{
		ID:          testCase.ID,
		CreatorId:   testCase.UserID,
		Title:       testCase.Title,
		Type:        testCase.Type,
		SuiteId:     testCase.SuiteID,
		Priority:    testCase.Priority,
		Description: testCase.Description,
		Steps:       testCase.Steps,
		Duration:    testCase.Duration,
		Status:      testCase.Status,
		CreatedAt:   testCase.CreatedAt,
		UpdatedAt:   testCase.UpdatedAt,
	}
}

func GetCaseResponses(testCases *[]models.Case) []CaseResponse {
	var cases []CaseResponse

	for _, testCase := range *testCases {
		cases = append(cases, GetCaseResponse(&testCase))
	}

	return cases
}
