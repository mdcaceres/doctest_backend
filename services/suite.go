package services

import (
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"strconv"
)

type ISuiteService interface {
	Create(payload *dto.SuiteRequest) (*dto.SuiteResponse, error)
	GetAll(projectId string) (*[]dto.SuiteResponse, error)
}

type SuiteService struct {
	SuiteProvider providers.SuiteProvider
}

func NewSuiteService() *SuiteService {
	return &SuiteService{
		SuiteProvider: providers.NewSuiteProvider(),
	}
}

func (s *SuiteService) Create(payload *dto.SuiteRequest) (*dto.SuiteResponse, error) {
	id, err := strconv.ParseUint(payload.ProjectId, 10, 64)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.ParseUint(payload.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	suite := &models.Suite{
		Name:        payload.Name,
		Description: payload.Description,
		Target:      payload.Target,
		ProjectID:   uint(id),
		UserId:      uint(userId),
	}

	suite, err = s.SuiteProvider.Create(suite)
	if err != nil {
		return nil, err
	}

	return dto.GetSuiteResponse(suite), nil
}

func (s *SuiteService) GetAll(projectId string) (*[]dto.SuiteResponse, error) {
	id, err := strconv.ParseUint(projectId, 10, 64)
	if err != nil {
		return nil, err
	}

	suites, err := s.SuiteProvider.GetAll(uint(id))
	if err != nil {
		return nil, err
	}

	return dto.GetSuiteResponses(suites), nil
}
