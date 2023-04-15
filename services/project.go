package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
)

type IProjectService interface {
	GetAll() (*[]dto.ProjectResponse, error)
	GetById() (*dto.ProjectResponse, error)
	Create() (*dto.ProjectResponse, error)
	Update() (*dto.ProjectResponse, error)
}

type ProjectService struct {
	ProjectProvider providers.ProjectProvider
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		ProjectProvider: providers.NewProjectProvider(),
	}
}

func (p *ProjectService) Create(c *fiber.Ctx, payload *dto.ProjectResponse) (*dto.ProjectResponse, error) {
	project := models.Project{
		Name:        payload.Name,
		Description: payload.Description,
		UserId:      payload.UserId,
	}

	createdProject, err := p.ProjectProvider.Create(&project)
	if err != nil {
		return nil, err
	}

	projectResponse := dto.GetProjectResponse(createdProject)

	return &projectResponse, nil
}
