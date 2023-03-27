package providers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
	"strings"
)

type UserProvider struct {
	DB *gorm.DB
}

func NewUserProvider() UserProvider {
	return UserProvider{
		DB: datasource.GetDB(),
	}
}

func (p *UserProvider) Create(ctx *fiber.Ctx, user *models.User) (*models.User, error) {
	result := p.DB.Create(user)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return nil, errors.New("user with that email already exists")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
