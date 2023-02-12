package dao

import "github.com/letscrum/letscrum/internal/model"

type WorkItemDao interface {
	Get(id int64) (*model.WorkItem, error)
	List(projectID int64, sprintID int64, featureID int64, page, size int32, keyword string) ([]*model.WorkItem, error)
	Count(projectID int64, sprintID int64, featureID int64, keyword string) int64
	Create(workItem *model.WorkItem) (int64, error)
	Update(workItem *model.WorkItem) (bool, error)
	Delete(id int64) (bool, error)
}
