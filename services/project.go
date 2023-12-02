package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"github.com/mdcaceres/doctest/services/mail"
	"os"
	"path"
	"strconv"
	"time"
)

type IProjectService interface {
	SaveProjectImage(projectID uint, fileBytes []byte) error
	Create(payload *dto.ProjectRequest) (*dto.ProjectResponse, error)
	Join(c *fiber.Ctx, payload *dto.JoinProject) (*dto.ProjectResponse, error)
	GetAll(userId string) (*[]dto.ProjectResponse, error)
	Get(id uint) (*dto.ProjectResponse, error)
}

type ProjectService struct {
	ProjectProvider providers.ProjectProvider
	UserProvider    providers.UserProvider
	EmailService    *mail.EmailService
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		ProjectProvider: providers.NewProjectProvider(),
		UserProvider:    providers.NewUserProvider(),
		EmailService:    mail.NewEmailService(),
	}
}

func (p *ProjectService) SaveProjectImage(projectID uint, fileBytes []byte) error {
	project, err := p.ProjectProvider.Get(&models.Project{ID: projectID})
	if err != nil {
		return err
	}

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%d_%s", project.ID, uuid.NewString())
	filePath := path.Join("uploads", fileName)

	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		return err
	}

	project.Image = filePath

	_, err = p.ProjectProvider.UpdateProject(project)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) Create(payload *dto.ProjectRequest) (*dto.ProjectResponse, error) {
	userId, err := strconv.ParseUint(payload.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	start, err := time.Parse(time.RFC3339Nano, payload.StartDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, payload.EndDate)
	if err != nil {
		return nil, err
	}

	project := models.Project{
		Name:            payload.Name,
		Description:     payload.Description,
		StartDate:       start,
		EndDate:         end,
		UserId:          uint(userId),
		ProjectClientID: payload.ClientId,
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

func (p *ProjectService) Get(id uint) (*dto.ProjectResponse, error) {
	project, err := p.ProjectProvider.Get(&models.Project{ID: id})

	if err != nil {
		return nil, err
	}

	response := dto.GetProjectResponse(project)

	return &response, nil
}
