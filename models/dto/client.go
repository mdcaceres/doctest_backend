package dto

import "github.com/mdcaceres/doctest/models"

type ProjectClientResponse struct {
	Id          uint              `json:"id"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	PhoneNumber string            `json:"phoneNumber"`
	Projects    []ProjectResponse `json:"projects"`
}

type ProjectClientRequest struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func GetProjectClientResponse(projectClient *models.ProjectClient) ProjectClientResponse {
	return ProjectClientResponse{
		Id:          projectClient.ID,
		Name:        projectClient.Name,
		Email:       projectClient.Email,
		PhoneNumber: projectClient.PhoneNumber,
	}
}

func GetProjectClientResponses(projectClient []models.ProjectClient) []ProjectClientResponse {
	var projects []ProjectClientResponse
	for _, project := range projectClient {
		projects = append(projects, GetProjectClientResponse(&project))
	}

	return projects
}
