package providers

import (
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type SuiteProvider struct {
	DB *gorm.DB
}

func NewSuiteProvider() SuiteProvider {
	return SuiteProvider{
		DB: datasource.GetDB(),
	}
}

func (s *SuiteProvider) Create(suite *models.Suite) (*models.Suite, error) {
	result := s.DB.Create(suite)
	if result.Error != nil {
		return nil, result.Error
	}
	return suite, nil
}

func (s *SuiteProvider) GetAll(projectId uint) (*[]models.Suite, error) {
	var suites []models.Suite

	result := s.DB.Find(&suites).Where("project_id = ?", projectId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &suites, nil
}
