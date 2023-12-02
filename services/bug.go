package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"github.com/mdcaceres/doctest/services/mail"
	"os"
	"path"
	"strconv"
	"strings"
)

type IBugService interface {
	Create(payload *dto.BugRequest) (*dto.BugResponse, error)
	GetAllByProjectId(projectID string) (*[]dto.BugResponse, error)
	GetAllByUserId(userId string, status string) (*[]dto.BugResponse, error)
	Update(payload *dto.BugRequest) (*dto.BugResponse, error)
	SaveFiles(bugId uint, files [][]byte) error
	AddComment(payload *dto.BugCommentRequest) (*dto.BugResponse, error)
}

type BugService struct {
	BugProvider  providers.BugProvider
	UserProvider providers.UserProvider
	EmailService *mail.EmailService
}

func NewBugService() *BugService {
	return &BugService{
		BugProvider:  providers.NewBugProvider(),
		UserProvider: providers.NewUserProvider(),
		EmailService: mail.NewEmailService(),
	}
}

func (b *BugService) Create(payload *dto.BugRequest) (*dto.BugResponse, error) {
	bug := payload.ToEntity()

	user, err := b.UserProvider.GetById(bug.UserID)

	users, err := b.UserProvider.GetByProject(bug.ProjectID)

	var mails []string

	for _, user := range *users {
		mails = append(mails, user.Email)
	}

	bug, err = b.BugProvider.Create(bug)
	if err != nil {
		return nil, err
	}

	mailData := dto.MailData{
		Name:    user.Name,
		To:      mails,
		Subject: "Your team mate has created a new bug in your project",
		Action:  "create new bug",
		Url:     fmt.Sprintf("%s%v", os.Getenv("VIEW_BUG"), bug.ID),
	}

	err = b.EmailService.SendSimple(&mailData)
	if err != nil {
		return nil, err
	}

	response := dto.GetBugResponse(bug)

	return &response, nil
}

func (b *BugService) GetAllByProjectId(projectID string) (*[]dto.BugResponse, error) {
	var responses []dto.BugResponse
	id, err := strconv.ParseUint(projectID, 10, 64)
	if err != nil {
		return nil, err
	}
	bugs, err := b.BugProvider.GetAllByProjectId(uint(id))
	if err != nil {
		return nil, err
	}

	responses = dto.GetBugResponses(*bugs)

	return &responses, nil
}

func (b *BugService) GetAllByUserId(userId string, status string) (*[]dto.BugResponse, error) {
	var responses []dto.BugResponse
	var filtered []models.Bug
	id, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, err
	}
	bugs, err := b.BugProvider.GetAllByUserId(uint(id))
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(*bugs); i++ {
		bug := (*bugs)[i]
		if strings.EqualFold(bug.Status, status) {
			filtered = append(filtered, bug)
		}
	}

	responses = dto.GetBugResponses(filtered)

	return &responses, nil
}

func (b *BugService) Update(payload *dto.BugRequest) (*dto.BugResponse, error) {
	bug := payload.ToEntity()
	user, err := b.UserProvider.GetById(bug.UserID)
	users, err := b.UserProvider.GetByProject(bug.ProjectID)
	var mails []string

	for _, user := range *users {
		mails = append(mails, user.Email)
	}

	bug = &models.Bug{
		ID:        payload.ID,
		Status:    payload.Status,
		UserID:    payload.UserID,
		ProjectID: payload.ProjectID,
	}

	bug, err = b.BugProvider.Update(bug)
	if err != nil {
		return nil, err
	}

	response := dto.GetBugResponse(bug)

	mailData := dto.MailData{
		Name:    user.Name,
		To:      mails,
		Subject: fmt.Sprintf("Your team mate has update the bug id %d to %s status", bug.ID, bug.Status),
		Action:  "update bug",
		Url:     fmt.Sprintf("%s%v", os.Getenv("VIEW_BUG"), bug.ID),
	}

	err = b.EmailService.SendSimple(&mailData)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (b *BugService) SaveFiles(bugId uint, files [][]byte) error {
	bug, err := b.BugProvider.GetById(bugId)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := fmt.Sprintf("%d_%s", bugId, uuid.NewString())
		filePath := path.Join("uploads", fileName)
		err := os.WriteFile(filePath, file, 0644)
		if err != nil {
			return err
		}
		bug.Files = append(bug.Files, filePath)
	}

	_, err = b.BugProvider.Update(bug)

	if err != nil {
		return err
	}

	return nil
}

func (b *BugService) AddComment(payload *dto.BugCommentRequest) (*dto.BugResponse, error) {
	comment := payload.ToEntity()

	bug, err := b.BugProvider.AddComment(comment)
	if err != nil {
		return nil, err
	}

	response := dto.GetBugResponse(bug)
	return &response, nil
}
