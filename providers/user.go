package providers

import (
	"errors"
	"fmt"
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

func (p *UserProvider) Create(user *models.User) (*models.User, error) {
	result := p.DB.Create(user)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return nil, errors.New("user with that email already exists")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (p *UserProvider) GetById(id uint) (*models.User, error) {
	user := models.User{
		ID: id,
	}
	result := p.DB.First(&user)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error in provider GetById ERROR: %s", result.Error))
	}
	return &user, nil
}

func (p *UserProvider) GetByName(username string) (*models.User, error) {
	user := models.User{
		Name: username,
	}
	result := p.DB.Where("name = ?", username).First(&user)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error in provider GetById ERROR: %s", result.Error))
	}
	return &user, nil
}

func (p *UserProvider) GetAll() (*[]models.User, error) {
	var users []models.User

	result := p.DB.Find(&users)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error in provider GetAll ERROR: %s", result.Error))
	}

	return &users, nil
}

func (p *UserProvider) GetByProject(projectId uint) (*[]models.User, error) {
	var users []models.User

	result := p.DB.Joins("JOIN project_team ON project_team.user_id = users.id").Where("project_team.project_id = ?", projectId).Find(&users)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error in provider GetAll ERROR: %s", result.Error))
	}

	return &users, nil
}

func (p *UserProvider) UpdateRoleById(id uint, roles []models.Role) (*models.User, error) {
	user := models.User{
		ID:    id,
		Roles: roles,
	}

	result := p.DB.Save(user)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error in provider Update Roles ERROR: %s", result.Error))
	}

	return &user, nil
}

func (p *UserProvider) UpdateFcmTokenById(id uint, token string) error {

	user := models.User{
		ID: id,
	}

	p.DB.First(&user)

	user.NotificationToken = token

	result := p.DB.Save(&user)

	if result.Error != nil {
		return errors.New(fmt.Sprintf("error in provider Update Roles ERROR: %s", result.Error))
	}

	return nil
}

func (p *UserProvider) UpdatePasswordById(id uint, password []byte) (*models.User, error) {
	user := models.User{
		ID:       id,
		Password: password,
	}

	result := p.DB.Save(user)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error in provider Update Password ERROR: %s", result.Error))
	}

	return &user, nil
}
