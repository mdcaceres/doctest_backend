package providers

import (
	"errors"
	"fmt"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"gorm.io/gorm"
)

type BugProvider struct {
	DB *gorm.DB
}

func NewBugProvider() BugProvider {
	return BugProvider{
		DB: datasource.GetDB(),
	}
}

func (b *BugProvider) Create(bug *models.Bug) (*models.Bug, error) {
	result := b.DB.Create(bug)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error creating new bug in our database [error:%v]", result.Error))
	}
	return bug, nil
}

func (b *BugProvider) GetAllByProjectId(projectId uint) (*[]models.Bug, error) {
	var bugs []models.Bug
	result := b.DB.Preload("Comments").Where("project_id = ?", projectId).Find(&bugs)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all bugs by project from our database [error:%v]", result.Error))
	}
	return &bugs, nil
}

func (b *BugProvider) GetById(id uint) (*models.Bug, error) {
	var bug models.Bug
	result := b.DB.First(&bug, id)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting bug by id from our database [error:%v]", result.Error))
	}
	return &bug, nil
}

func (b *BugProvider) GetAllByUserId(userId uint) (*[]models.Bug, error) {
	var bugs []models.Bug
	result := b.DB.Where("user_id = ?", userId).Find(&bugs)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error getting all bugs bu user from our database [error:%v]", result.Error))
	}
	return &bugs, nil
}

func (b *BugProvider) Update(bug *models.Bug) (*models.Bug, error) {
	result := b.DB.Model(&models.Bug{}).Where("id = ?", bug.ID).Update("status", bug.Status)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error updating bug files in our database [error:%v]", result.Error))
	}
	return bug, nil
}

func (b *BugProvider) AddComment(bugComment *models.BugComment) (*models.Bug, error) {

	bug, err := b.GetById(bugComment.BugID)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error adding comment to bug in our database [error:%v]", err))
	}

	bug.Comments = append(bug.Comments, *bugComment)

	result := datasource.GetDB().Save(&bug)

	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("error adding comment to bug in our database [error:%v]", result.Error))
	}

	return bug, nil
}
