package dto

import (
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/media"
	"time"
)

type BugRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	ProjectID   uint   `json:"project_id"`
	//TestCaseID  uint   `json:"test_case_id"`
	Status     string `json:"status"`
	Priority   string `json:"priority"`
	Severity   string `json:"severity"`
	AssignedID uint   `json:"assigned_id"`
	Due        string `json:"due"`
}

type BugResponse struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	UserId      uint                 `json:"user_id"`
	ProjectID   uint                 `json:"project_id"`
	TestCaseID  uint                 `json:"test_case_id"`
	Status      string               `json:"status"`
	Priority    string               `json:"priority"`
	Severity    string               `json:"severity"`
	AssignedId  uint                 `json:"assigned_id"`
	Files       media.Files          `gorm:"type:VARCHAR(255)"`
	CreatedAt   time.Time            `json:"created_at"`
	Due         time.Time            `json:"due"`
	Comments    []BugCommentResponse `json:"comments"`
}

func GetBugResponse(bug *models.Bug) BugResponse {
	var bugComments []BugCommentResponse

	for _, comment := range bug.Comments {
		bugComments = append(bugComments, GetBugCommentResponse(&comment))
	}
	return BugResponse{
		ID:          bug.ID,
		Name:        bug.Name,
		Description: bug.Description,
		UserId:      bug.UserID,
		ProjectID:   bug.ProjectID,
		//TestCaseID:  bug.TestCaseID,
		Status:     bug.Status,
		Priority:   bug.Priority,
		Severity:   bug.Severity,
		AssignedId: bug.AssignedId,
		Files:      bug.Files,
		CreatedAt:  bug.CreatedAt,
		Due:        bug.Due,
		Comments:   bugComments,
	}
}

func GetBugResponses(bugs []models.Bug) []BugResponse {
	var responses []BugResponse
	for _, bug := range bugs {
		responses = append(responses, GetBugResponse(&bug))
	}
	return responses
}

func (b *BugRequest) ToEntity() *models.Bug {
	return &models.Bug{
		ID:          b.ID,
		Name:        b.Name,
		Description: b.Description,
		UserID:      b.UserID,
		ProjectID:   b.ProjectID,
		//TestCaseID:  b.TestCaseID,
		Status:     b.Status,
		Priority:   b.Priority,
		Severity:   b.Severity,
		AssignedId: b.AssignedID,
		Due:        time.Now(),
	}
}
