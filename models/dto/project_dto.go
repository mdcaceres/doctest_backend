package dto

import "github.com/mdcaceres/doctest/models"

type ProjectResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
}

func GetProjectResponse(project *models.Project) ProjectResponse {
	return ProjectResponse{
		Id:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		UserId:      project.UserId,
	}
}
