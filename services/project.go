package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"strconv"
)

type IProjectService interface {
	Join(c *fiber.Ctx, payload *dto.JoinProject)
	Create(c *fiber.Ctx, payload *dto.ProjectResponse) (*dto.ProjectResponse, error)
}

type ProjectService struct {
	ProjectProvider providers.ProjectProvider
	UserProvider    providers.UserProvider
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		ProjectProvider: providers.NewProjectProvider(),
		UserProvider:    providers.NewUserProvider(),
	}
}

func (p *ProjectService) Create(c *fiber.Ctx, payload *dto.ProjectResponse) (*dto.ProjectResponse, error) {
	userId, err := strconv.ParseUint(payload.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	project := models.Project{
		Name:        payload.Name,
		Description: payload.Description,
		UserId:      uint(userId),
	}

	createdProject, err := p.ProjectProvider.Create(&project)
	if err != nil {
		return nil, err
	}

	projectResponse := dto.GetProjectResponse(createdProject)

	return &projectResponse, nil
}

func (p *ProjectService) Join(c *fiber.Ctx, payload *dto.JoinProject) (*dto.ProjectResponse, error) {
	userId, err := strconv.ParseUint(payload.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	projectId, err := strconv.ParseUint(payload.ProjectId, 10, 64)
	if err != nil {
		return nil, err
	}

	project := models.Project{
		ID: uint(projectId),
	}

	p.ProjectProvider.Get(&project)

	if err != nil {
		return nil, err
	}

	user, err := p.UserProvider.GetById(uint(userId))
	if err != nil {
		return nil, err
	}

	team := append(project.Team, user)

	project.Team = team

	createdProject, err := p.ProjectProvider.UpdateProject(&project)
	if err != nil {
		return nil, err
	}

	response := dto.GetProjectResponse(createdProject)

	return &response, nil
}

func (p *ProjectService) GetAll(userId string) (*[]dto.ProjectResponse, error) {
	var responses []dto.ProjectResponse

	projects, err := p.ProjectProvider.GetAll(userId)

	if err != nil {
		return nil, err
	}

	for _, p := range *projects {
		responses = append(responses, dto.GetProjectResponse(&p))
	}

	return &responses, nil
}
