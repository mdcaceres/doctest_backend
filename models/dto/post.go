package dto

import (
	"github.com/mdcaceres/doctest/models"
	"time"
)

type PostRequest struct {
	ID        uint   `json:"id"`
	Comment   string `json:"comment"`
	ProjectID uint   `json:"project_id"`
	UserID    uint   `json:"user_id"`
	Image     string
}

type PostResponse struct {
	ID        uint   `json:"id"`
	Comment   string `json:"comment"`
	ProjectID uint   `json:"project_id"`
	UserID    uint   `json:"user_id"`
	Image     string
}

func GetPostResponse(p *models.Post) PostResponse {
	return PostResponse{
		ID:        p.ID,
		UserID:    p.UserID,
		ProjectID: p.ProjectID,
		Comment:   p.Comment,
		Image:     p.Image,
	}
}

func GetPostResponses(Posts []models.Post) []PostResponse {
	var responses []PostResponse
	for _, Post := range Posts {
		responses = append(responses, GetPostResponse(&Post))
	}
	return responses
}

func (p *PostRequest) ToEntity() *models.Post {
	return &models.Post{
		ID:        p.ID,
		CreatedAt: time.Now(),
		Comment:   p.Comment,
		ProjectID: p.ProjectID,
		UserID:    p.UserID,
		Image:     p.Image,
	}
}
