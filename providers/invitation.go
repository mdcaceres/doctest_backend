package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type InvitationProvider struct {
	DB *gorm.DB
}

func NewInvitationProvider() InvitationProvider {
	return InvitationProvider{
		DB: datasource.GetDB(),
	}
}

func (p *InvitationProvider) Create(invitation *models.Invitation) (*models.Invitation, error) {
	result := p.DB.Create(invitation)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error creating new invitation in our database [error:%v]", result.Error))
	}
	return invitation, nil
}

func (p *InvitationProvider) UpdateInvitation(invitation *models.Invitation) (*models.Invitation, error) {
	result := p.DB.Save(invitation)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error updating invitation [error:%v]", result.Error))
	}
	return invitation, nil
}

func (p *InvitationProvider) Get(invitation *models.Invitation) (*models.Invitation, error) {
	result := p.DB.First(invitation)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("invitation id : %v not exists", invitation.ID))
	}
	return invitation, nil
}

func (p *InvitationProvider) GeAllByInvitedId(invitation *models.Invitation) ([]*models.Invitation, error) {
	var invitations []*models.Invitation
	result := p.DB.Where("invited_id = ?", invitation.InvitedID).Find(&invitations)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("invitations with invited_id : %v not found", invitation.InvitedID))
	}
	return invitations, nil
}
