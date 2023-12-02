package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type PostProvider struct {
	DB *gorm.DB
}

func NewPostProvider() PostProvider {
	return PostProvider{
		DB: datasource.GetDB(),
	}
}

func (b *PostProvider) Create(Post *models.Post) (*models.Post, error) {
	result := b.DB.Create(Post)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error creating new Post in our database [error:%v]", result.Error))
	}
	return Post, nil
}

func (b *PostProvider) GetAllByProjectId(projectId uint) (*[]models.Post, error) {
	var Posts []models.Post
	result := b.DB.Where("project_id = ?", projectId).Find(&Posts)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all Posts by project from our database [error:%v]", result.Error))
	}
	return &Posts, nil
}

func (b *PostProvider) GetById(id uint) (*models.Post, error) {
	var Post models.Post
	result := b.DB.First(&Post, id)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting Post by id from our database [error:%v]", result.Error))
	}
	return &Post, nil
}

func (b *PostProvider) GetAllByUserId(userId uint) (*[]models.Post, error) {
	var Posts []models.Post
	result := b.DB.Where("user_id = ?", userId).Find(&Posts)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all Posts bu user from our database [error:%v]", result.Error))
	}
	return &Posts, nil
}
