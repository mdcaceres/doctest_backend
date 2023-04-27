package dto

import "github.com/mdcaceres/doctest/models"

type ProjectResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      string `json:"userId"`
}

type JoinProject struct {
	ProjectId string `json:"projectId"`
	UserId    string `json:"userId"`
}

func GetProjectResponse(project *models.Project) ProjectResponse {
	return ProjectResponse{
		Id:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		UserId:      string(project.UserId),
	}
}
