package dto

import "github.com/mdcaceres/doctest/models"

type InvitationRequest struct {
	InviterID string `json:"inviterId"`
	InvitedID string `json:"invitedId"`
	ProjectId string `json:"projectId"`
}

type InvitationResponse struct {
	InviterID string `json:"inviterId"`
	InvitedID string `json:"invitedId"`
	ProjectId string `json:"projectId"`
	Accepted  bool   `gorm:"default:false"`
}

func GetInvitationResponse(invitation *models.Invitation) InvitationResponse {
	return InvitationResponse{
		InviterID: string(invitation.InviterID),
		InvitedID: string(invitation.InvitedID),
		ProjectId: string(invitation.ProjectID),
		Accepted:  invitation.Accepted,
	}
}
