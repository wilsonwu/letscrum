package service

import (
	"context"
	generalv1 "github.com/letscrum/letscrum/api/general/v1"
	letscrumv1 "github.com/letscrum/letscrum/api/letscrum/v1"
)

type DemoService struct {
	// This is generated by protoc
	letscrumv1.UnimplementedDemoServer
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (s *DemoService) Demo(ctx context.Context, req *generalv1.DemoRequest) (*generalv1.DemoResponse, error) {
	return &generalv1.DemoResponse{
		Demo: &generalv1.Demo{
			Demo: req.Demo,
		},
	}, nil
}