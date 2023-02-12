package mysql

import (
	"github.com/letscrum/letscrum/internal/model"
	"gorm.io/gorm"
)

type WorkItemDao struct {
	DB *gorm.DB
}

func (w WorkItemDao) Get(id int64) (*model.WorkItem, error) {
	var workItem *model.WorkItem
	query := w.DB.Model(&model.WorkItem{}).Where("id = ?", id)
	err := query.Find(&workItem).Error
	if err != nil {
		return nil, err
	}
	return workItem, nil
}

func (w WorkItemDao) List(projectID int64, sprintID int64, featureID int64, page, size int32, keyword string) ([]*model.WorkItem, error) {
	var workItems []*model.WorkItem
	query := w.DB.Model(&model.WorkItem{}).Where("project_id = ?", projectID).Where("title LIKE ?", "%"+keyword+"%").Limit(int(size)).Offset(int((page - 1) * size))
	if sprintID > 0 {
		query = query.Where("sprint_id", sprintID)
	}
	if featureID > 0 {
		query = query.Where("feature_id", featureID)
	}
	err := query.Find(&workItems).Error
	if err != nil {
		return nil, err
	}
	return workItems, nil
}

func (w WorkItemDao) Count(projectID int64, sprintID int64, featureID int64, keyword string) int64 {
	count := int64(0)
	query := w.DB.Model(&model.WorkItem{}).Where("project_id = ?", projectID).Where("sprint_id", sprintID).Where("title LIKE ?", "%"+keyword+"%")
	if sprintID > 0 {
		query = query.Where("sprint_id", sprintID)
	}
	if featureID > 0 {
		query = query.Where("feature_id", featureID)
	}
	query.Count(&count)
	return count
}

func (w WorkItemDao) Create(workItem *model.WorkItem) (int64, error) {
	if err := w.DB.Create(&workItem).Error; err != nil {
		return 0, err
	}
	return workItem.ID, nil
}

func (w WorkItemDao) Update(workItem *model.WorkItem) (bool, error) {
	if err := w.DB.Model(&model.WorkItem{}).Where("id = ?", workItem.ID).Update("title", workItem.Title).Error; err != nil {
		return false, nil
	}
	return true, nil
}

func (w WorkItemDao) Delete(id int64) (bool, error) {
	if err := w.DB.Where("id = ?", id).Delete(&model.WorkItem{}).Error; err != nil {
		return false, nil
	}
	return true, nil
}

func NewWorkItemDao(d *gorm.DB) *WorkItemDao {
	return &WorkItemDao{d}
}
