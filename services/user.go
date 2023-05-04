package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/auth"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type IUserService interface {
	GetAll() (*[]dto.UserResponse, error)
	GetById(id int64) (*dto.UserResponse, error)
	Create(user auth.SignUpInput) (*dto.UserResponse, error)
	UpdateRole(roles models.Role) (*dto.UserResponse, error)
}

type UserService struct {
	UserProvider providers.UserProvider
}

func NewUserService() *UserService {
	return &UserService{
		UserProvider: providers.NewUserProvider(),
	}
}

func (u *UserService) GetAll() (*[]dto.UserResponse, error) {
	var responses []dto.UserResponse

	users, err := u.UserProvider.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		responses = append(responses, dto.GetUserResponse(&user))
	}

	return &responses, nil
}

func (u *UserService) GetById(id uint) (*dto.UserResponse, error) {
	user, err := u.UserProvider.GetById(id)
	if err != nil {
		return nil, err
	}

	userResponse := dto.GetUserResponse(user)

	return &userResponse, nil
}

func (u *UserService) GetByUsername(username string) (*dto.UserResponse, error) {
	user, err := u.UserProvider.GetByName(username)

	if err != nil {
		return nil, err
	}

	userResponse := dto.GetUserResponse(user)

	return &userResponse, nil
}

func (u *UserService) Create(c *fiber.Ctx, payload *auth.SignUpInput) (*dto.UserResponse, error) {
	if payload.Password != payload.Confirm {
		return nil, errors.New("password do not match")
	}

	encryptPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	role := models.Role{Name: models.Tester}

	user := models.User{
		Name:              payload.Name,
		Email:             strings.ToLower(payload.Email),
		Photo:             payload.Photo,
		Password:          encryptPass,
		NotificationToken: payload.NotificationToken,
		Roles:             []models.Role{role},
	}

	createdUser, err := u.UserProvider.Create(&user)

	if err != nil {
		return nil, err
	}

	userResponse := dto.GetUserResponse(createdUser)

	return &userResponse, nil
}

func (u *UserService) UpdateRole(id uint, roles []models.Role) (*dto.UserResponse, error) {
	updatedUser, err := u.UserProvider.UpdateRoleById(id, roles)

	if err != nil {
		return nil, err
	}

	userResponse := dto.GetUserResponse(updatedUser)

	return &userResponse, nil
}

func (u *UserService) UpdateFcmToken(id uint, token string) error {
	err := u.UserProvider.UpdateFcmTokenById(id, token)

	if err != nil {
		errors.New(fmt.Sprintf("error updating token %v", err.Error()))
	}

	return nil
}
