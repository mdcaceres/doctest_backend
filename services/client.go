package services

import (
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"strconv"
)

type IClientService interface {
	Create(payload *dto.ProjectClientRequest) (*dto.ProjectClientResponse, error)
}

type ClientService struct {
	ClientProvider providers.ClientProvider
}

func NewClientService() *ClientService {
	return &ClientService{
		ClientProvider: providers.NewClientProvider(),
	}
}

func (c *ClientService) Create(payload *dto.ProjectClientRequest) (*dto.ProjectClientResponse, error) {
	client := &models.ProjectClient{
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
	}

	_, err := c.ClientProvider.CreateClient(client)

	if err != nil {
		return nil, err
	}

	resp := dto.GetProjectClientResponse(client)

	return &resp, nil
}

func (c *ClientService) GetAll(userId string) (*[]dto.ProjectClientResponse, error) {
	id, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	clients, err := c.ClientProvider.GetAll(id)

	if err != nil {
		return nil, err
	}

	resp := dto.GetProjectClientResponses(*clients)

	return &resp, nil
}
