package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type CaseProvider struct {
	DB *gorm.DB
}

func NewCaseProvider() CaseProvider {
	return CaseProvider{
		DB: datasource.GetDB(),
	}
}

func (c *CaseProvider) Create(testCase *models.Case) (*models.Case, error) {
	result := c.DB.Create(testCase)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error creating new test case in our database [error:%v]", result.Error))
	}
	return testCase, nil
}

func (c *CaseProvider) GetAllByProjectId(projectId uint) (*[]models.Case, error) {
	var cases []models.Case

	result := c.DB.Find(&cases).Where("project_id = ?", projectId)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all test cases from our database [error:%v]", result.Error))
	}

	return &cases, nil
}

func (c *CaseProvider) Get(testCase *models.Case) (*models.Case, error) {
	result := c.DB.First(testCase)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("test case id : %v not exists", testCase.ID))
	}
	return testCase, nil
}

func (c *CaseProvider) GetAllBySuiteId(suiteId uint) (*[]models.Case, error) {
	var cases []models.Case

	result := c.DB.Find(&cases).Where("suite_id = ?", suiteId)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all test cases from our database [error:%v]", result.Error))
	}

	return &cases, nil
}

func (c *CaseProvider) UpdateCaseFiles(test *models.Case) (*models.Case, error) {
	result := c.DB.Save(test)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error updating test case [error:%v]", result.Error))
	}
	return test, nil
}
