package dto

import (
	"github.com/mdcaceres/doctest/models"
	"time"
)

type UserResponse struct {
	Id        uint          `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Email     string        `json:"email,omitempty"`
	Role      []models.Role `json:"role,omitempty"`
	Photo     *string       `json:"photo,omitempty"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
}

func GetUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Roles,
		Photo:     user.Photo,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
