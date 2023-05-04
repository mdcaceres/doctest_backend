package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type ProjectProvider struct {
	DB *gorm.DB
}

func NewProjectProvider() ProjectProvider {
	return ProjectProvider{
		DB: datasource.GetDB(),
	}
}

func (p *ProjectProvider) Create(project *models.Project) (*models.Project, error) {
	result := p.DB.Create(project)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error creating new project in our database [error:%v]", result.Error))
	}
	return project, nil
}

func (p *ProjectProvider) UpdateProject(project *models.Project) (*models.Project, error) {
	result := p.DB.Save(project)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error updating project [error:%v]", result.Error))
	}
	return project, nil
}

func (p *ProjectProvider) Get(project *models.Project) (*models.Project, error) {
	result := p.DB.First(project)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("project id : %v not exists", project.ID))
	}
	return project, nil
}

func (p *ProjectProvider) GetAll(userId string) (*[]models.Project, error) {
	var projects []models.Project

	result := p.DB.Find(&projects).Preload("team", "user_id = ?", userId)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error in project provider GetAll ERROR: %s", result.Error))
	}

	return &projects, nil
}
