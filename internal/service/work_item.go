package service

import (
	"context"
	generalv1 "github.com/letscrum/letscrum/api/general/v1"
	itemv1 "github.com/letscrum/letscrum/api/item/v1"
	v1 "github.com/letscrum/letscrum/api/letscrum/v1"
	userV1 "github.com/letscrum/letscrum/api/user/v1"
	"github.com/letscrum/letscrum/internal/dao"
	"github.com/letscrum/letscrum/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkItemService struct {
	v1.UnimplementedWorkItemServer
	workItemDao dao.WorkItemDao
}

func NewWorkItemService(dao dao.Interface) *WorkItemService {
	return &WorkItemService{
		workItemDao: dao.WorkItemDao(),
	}
}

func (s *WorkItemService) List(ctx context.Context, req *itemv1.ListWorkItemRequest) (*itemv1.ListWorkItemResponse, error) {
	_, err := utils.AuthJWT(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	workItems, err := s.workItemDao.List(req.ProjectId, req.SprintId, req.FeatureId, req.Page, req.Size, req.Keyword)
	if err != nil {
		result := status.Convert(err)
		if result.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "not found.")
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}
	var list []*itemv1.WorkItem
	for _, w := range workItems {
		var workItem = &itemv1.WorkItem{
			Id:        w.ID,
			ProjectId: w.ProjectID,
			SprintId:  w.SprintID,
			FeatureId: w.FeatureID,
			Title:     w.Title,
			AssignUser: &userV1.User{
				Id:   w.AssignUser.ID,
				Name: w.AssignUser.Name,
			},
			CreatedUser: &userV1.User{
				Id:   w.CreatedUser.ID,
				Name: w.CreatedUser.Name,
			},
			CreatedAt: w.CreatedAt.Unix(),
			UpdatedAt: w.UpdatedAt.Unix(),
		}
		list = append(list, workItem)
	}
	count := s.workItemDao.Count(req.ProjectId, req.SprintId, req.FeatureId, req.Keyword)
	return &itemv1.ListWorkItemResponse{
		Items: list,
		Pagination: &generalv1.Pagination{
			Page:  req.Page,
			Size:  req.Size,
			Total: int32(count),
		},
	}, nil
}
