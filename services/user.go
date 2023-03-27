package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/auth"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type IUserService interface {
	GetAll() []models.User
	GetByID(id int64) models.User
	GetByName(name string) models.User
	Create(user auth.SignUpInput) models.User
}

type UserService struct {
	UserProvider providers.UserProvider
}

func NewUserService() *UserService {
	return &UserService{
		UserProvider: providers.NewUserProvider(),
	}
}

func (u *UserService) GetAll() []models.User {
	//todo
	return nil
}

func (u *UserService) GetById(id int64) models.User {
	//todo
	return models.User{}
}

func (u *UserService) GetByName(name string) models.User {
	//todo
	return models.User{}
}

func (u *UserService) Create(c *fiber.Ctx, payload *auth.SignUpInput) (*dto.UserResponse, error) {
	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("password do not match")
	}

	encryptPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Photo:    &payload.Photo,
		Password: encryptPass,
	}

	createdUser, err := u.UserProvider.Create(c, &user)

	if err != nil {
		return nil, errors.New("error when provider creates a user")
	}

	userResponse := dto.GetUserResponse(createdUser)

	return &userResponse, nil
}
