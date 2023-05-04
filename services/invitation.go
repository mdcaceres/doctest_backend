package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"strconv"
)

type IInvitationService interface {
	Create(c *fiber.Ctx, payload *dto.InvitationRequest) (*dto.InvitationResponse, error)
	Get(c *fiber.Ctx, id uint) (*dto.InvitationResponse, error)
	Update(c *fiber.Ctx, id uint, payload *dto.InvitationRequest) (*dto.InvitationResponse, error)
}

type InvitationService struct {
	InvitationProvider providers.InvitationProvider
	ProjectProvider    providers.ProjectProvider
	UserProvider       providers.UserProvider
	MessagingProvider  providers.MessagingProvider
}

func NewInvitationService() *InvitationService {
	return &InvitationService{
		InvitationProvider: providers.NewInvitationProvider(),
		UserProvider:       providers.NewUserProvider(),
		ProjectProvider:    providers.NewProjectProvider(),
		MessagingProvider:  providers.NewMessagingProvider(),
	}
}

func (i *InvitationService) Create(c *fiber.Ctx, payload *dto.InvitationRequest) (*dto.InvitationResponse, error) {
	projectId, err := strconv.ParseUint(payload.ProjectId, 10, 64)
	if err != nil {
		return nil, err
	}

	invitedId, err := strconv.ParseUint(payload.InvitedID, 10, 64)
	if err != nil {
		return nil, err
	}

	invited, err := i.UserProvider.GetById(uint(invitedId))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user %d not found", invitedId))
	}

	inviterId, err := strconv.ParseUint(payload.InviterID, 10, 64)
	if err != nil {
		return nil, err
	}

	inviter, err := i.UserProvider.GetById(uint(inviterId))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user %d not found", inviter.ID))
	}

	role := models.GetRoleByName(payload.Role)

	_, err = i.ProjectProvider.Get(
		&models.Project{
			ID: uint(projectId),
		})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("project %d not found", projectId))
	}

	/*
		_, err = i.InvitationProvider.Get(&models.Invitation{ProjectID: uint(projectId), InvitedID: uint(invitedId)})
		if err == nil {
			return nil, errors.New(fmt.Sprintf("invitation for email %s already exists in project %d", payload.InviterID, projectId))
		}
	*/
	invitation := models.Invitation{
		ProjectID:    uint(projectId),
		InviterID:    inviter.ID,
		InvitedID:    invited.ID,
		Role:         role.Name,
		InvitedEmail: invited.Email,
		Accepted:     false,
	}

	createdInvitation, err := i.InvitationProvider.Create(&invitation)

	if err != nil {
		return nil, err
	}

	//i.MessagingProvider.SendToToken(invited.NotificationToken)

	response := dto.GetInvitationResponse(createdInvitation)

	return &response, nil
}

func (i *InvitationService) GetAllByInvitedId(c *fiber.Ctx, invitedId uint) (*[]*dto.InvitationResponse, error) {
	invitation := &models.Invitation{InvitedID: invitedId}

	invitations, err := i.InvitationProvider.GeAllByInvitedId(invitation)
	if err != nil {
		return nil, err
	}

	var response []*dto.InvitationResponse

	for _, invitation := range invitations {
		dto := dto.GetInvitationResponse(invitation)
		response = append(response, &dto)
	}

	return &response, nil
}

func (i *InvitationService) Update(c *fiber.Ctx, id uint) (*dto.InvitationResponse, error) {
	invitation := &models.Invitation{ID: id}

	invitation, err := i.InvitationProvider.Get(invitation)

	if err != nil {
		return nil, err
	}

	invitation.Accepted = true

	updatedInvitation, err := i.InvitationProvider.UpdateInvitation(invitation)
	if err != nil {
		return nil, err
	}

	response := dto.GetInvitationResponse(updatedInvitation)
	return &response, nil
}
