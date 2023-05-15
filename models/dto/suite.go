package dto

import "github.com/mdcaceres/doctest/models"

type SuiteRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Target      string `json:"target"`
	ProjectId   string `json:"project_id"`
	UserId      string `json:"user_id"`
}

type SuiteResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Target      string `json:"target"`
}

func GetSuiteResponse(suite *models.Suite) *SuiteResponse {
	return &SuiteResponse{
		Id:          suite.ID,
		Name:        suite.Name,
		Description: suite.Description,
		Target:      suite.Target,
	}
}

func GetSuiteResponses(suites *[]models.Suite) *[]SuiteResponse {
	var suiteResponses []SuiteResponse

	for _, suite := range *suites {
		suiteResponses = append(suiteResponses, *GetSuiteResponse(&suite))
	}

	return &suiteResponses
}
