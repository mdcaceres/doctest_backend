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
