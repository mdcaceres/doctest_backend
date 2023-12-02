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
		return nil, errors.New(fmt.Sprintf("error creating new TestExecution case in our database [error:%v]", result.Error))
	}
	return testCase, nil
}

func (c *CaseProvider) GetAllByProjectId(projectId uint) (*[]models.Case, error) {
	var cases []models.Case
	result := c.DB.Where("project_id = ?", projectId).Find(&cases)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all TestExecution cases from our database [error:%v]", result.Error))
	}
	return &cases, nil
}

func (c *CaseProvider) GetAllByUserId(userId uint) (*[]models.Case, error) {
	var cases []models.Case
	result := c.DB.Where("user_id = ?", userId).Order("created_at desc").Find(&cases)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all TestExecution cases from our database [error:%v]", result.Error))
	}
	return &cases, nil
}

func (c *CaseProvider) Get(testCase *models.Case) (*models.Case, error) {
	result := c.DB.Preload("Steps").First(testCase)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("TestExecution case id : %v not exists", testCase.ID))
	}
	return testCase, nil
}

func (c *CaseProvider) GetAllBySuiteId(suiteId uint) (*[]models.Case, error) {
	var cases []models.Case

	result := c.DB.Find(&cases).Where("suite_id = ?", suiteId)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all TestExecution cases from our database [error:%v]", result.Error))
	}

	return &cases, nil
}

func (c *CaseProvider) Update(test *models.Case) (*models.Case, error) {
	result := c.DB.Save(test)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error updating TestExecution case [error:%v]", result.Error))
	}
	return test, nil
}
