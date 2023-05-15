package dto

import (
	"github.com/mdcaceres/doctest/models"
	"strconv"
	"time"
)

type ProjectResponse struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	UserId      string    `json:"userId"`
	Image       string    `json:"image"`
}

type ProjectRequest struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	UserId      string `json:"userId"`
	ClientId    string `json:"clientId"`
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
		StartDate:   project.StartDate,
		EndDate:     project.EndDate,
		UserId:      strconv.FormatUint(uint64(project.UserId), 10),
		Image:       project.Image,
	}
}
