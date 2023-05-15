package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type ClientProvider struct {
	DB *gorm.DB
}

func NewClientProvider() ClientProvider {
	return ClientProvider{
		DB: datasource.GetDB(),
	}
}

func (c *ClientProvider) CreateClient(client *models.ProjectClient) (*models.ProjectClient, error) {
	result := c.DB.Create(client)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error creating new client in our database [error:%v]", result.Error))
	}
	return client, nil
}
