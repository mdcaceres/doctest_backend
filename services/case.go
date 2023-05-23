package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/dto"
	"github.com/mdcaceres/doctest/providers"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type ICaseService interface {
	Create(payload *dto.CaseRequest) (*dto.CaseResponse, error)
	GetAllByProjectId(projectID uint) (*[]dto.CaseResponse, error)
	GetAllBySuiteId(suiteID uint) (*[]dto.CaseResponse, error)
}

type CaseService struct {
	CaseProvider providers.CaseProvider
}

func NewCaseService() *CaseService {
	return &CaseService{
		CaseProvider: providers.NewCaseProvider(),
	}
}

func (c *CaseService) Create(payload *dto.CaseRequest) (*dto.CaseResponse, error) {
	creatorId, err := strconv.ParseInt(payload.CreatorId, 10, 64)
	if err != nil {
		return nil, err
	}
	projectId, err := strconv.ParseInt(payload.ProjectId, 10, 64)
	if err != nil {
		return nil, err
	}
	/*suiteId, err := strconv.ParseInt(payload.SuiteId, 10, 64)
	if err != nil {
		return nil, err
	}*/
	duration, err := time.ParseDuration(strings.Replace(payload.Duration, ":", "m", 1) + "s")
	if err != nil {
		return nil, err
	}
	steps := make([]models.Step, 0)
	order := uint(1)

	for _, step := range payload.Steps {
		newStep := models.Step{
			Order:       order,
			Description: step["description"],
			Result:      step["result"]}
		steps = append(steps, newStep)
		order++
	}

	caseModel := models.Case{
		Title:       payload.Title,
		UserID:      uint(creatorId),
		Type:        payload.Type,
		Priority:    payload.Priority,
		Description: payload.Description,
		Duration:    duration,
		Steps:       steps,
		ProjectId:   uint(projectId),
		SuiteId:     payload.SuiteId,
		Status:      payload.Status,
	}

	created, err := c.CaseProvider.Create(&caseModel)

	if err != nil {
		return nil, err
	}
	caseResponse := dto.GetCaseResponse(created)

	return &caseResponse, nil
}

func (c *CaseService) SaveFiles(caseId uint, files [][]byte) error {
	test := &models.Case{
		ID: caseId,
	}

	t, err := c.CaseProvider.Get(test)
	if err != nil {
		return err
	}

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := fmt.Sprintf("%d_%s", caseId, uuid.NewString())
		filePath := path.Join("uploads", fileName)
		err = os.WriteFile(filePath, file, 0644)
		if err != nil {
			return err
		}
		t.Files = append(t.Files, filePath)
	}

	_, err = c.CaseProvider.Update(t)

	if err != nil {
		return err
	}

	return nil
}

func (c *CaseService) GetAllByProjectId(id string) (*[]dto.CaseResponse, error) {
	projectId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	cases, err := c.CaseProvider.GetAllByProjectId(uint(projectId))

	if err != nil {
		return nil, err
	}

	casesResponse := dto.GetCaseResponses(cases)

	return &casesResponse, nil
}
