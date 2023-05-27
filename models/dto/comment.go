package dto

import "github.com/mdcaceres/doctest/models"

type TestCommentRequest struct {
	Comment string `json:"comment"`
	CaseID  uint   `json:"test_id"`
	UserID  uint   `json:"user_id"`
}

type BugCommentRequest struct {
	Comment string `json:"comment"`
	BugID   uint   `json:"bug_id"`
	UserID  uint   `json:"user_id"`
}

type BugCommentResponse struct {
	ID          uint   `json:"id"`
	DateCreated string `json:"date_created"`
	Comment     string `json:"comment"`
	BugID       uint   `json:"bug_id"`
	UserID      uint   `json:"user_id"`
}

type TestCommentResponse struct {
	ID          uint   `json:"id"`
	DateCreated string `json:"date_created"`
	Comment     string `json:"comment"`
	CaseID      uint   `json:"test_id"`
	UserID      uint   `json:"user_id"`
}

func (b *BugCommentRequest) ToEntity() *models.BugComment {
	return &models.BugComment{
		Comment: b.Comment,
		BugID:   b.BugID,
		UserID:  b.UserID,
	}
}

func GetBugCommentResponse(c *models.BugComment) BugCommentResponse {
	return BugCommentResponse{
		ID:          c.ID,
		DateCreated: c.CreatedAt.String(),
		Comment:     c.Comment,
		BugID:       c.BugID,
		UserID:      c.UserID,
	}
}
