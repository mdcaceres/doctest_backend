package services

import (
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"github.com/mdcaceres/doctest/services/mail"
	"strconv"
)

type IPostService interface {
	Create(payload *dto.PostRequest) (*dto.PostResponse, error)
	GetAllByProjectId(projectID string) (*[]dto.PostResponse, error)
}

type PostService struct {
	PostProvider providers.PostProvider
	UserProvider providers.UserProvider
	EmailService *mail.EmailService
}

func NewPostService() *PostService {
	return &PostService{
		PostProvider: providers.NewPostProvider(),
		UserProvider: providers.NewUserProvider(),
		EmailService: mail.NewEmailService(),
	}
}

func (b *PostService) Create(payload *dto.PostRequest) (*dto.PostResponse, error) {
	post := payload.ToEntity()

	p, e := b.PostProvider.Create(post)
	if e != nil {
		return nil, e
	}

	response := dto.GetPostResponse(p)

	return &response, nil
}

func (b *PostService) GetAllByProjectId(projectID string) (*[]dto.PostResponse, error) {
	var responses []dto.PostResponse
	id, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		return nil, err
	}
	p, err := b.PostProvider.GetAllByProjectId(uint(id))
	if err != nil {
		return nil, err
	}

	responses = dto.GetPostResponses(*p)

	return &responses, nil
}
