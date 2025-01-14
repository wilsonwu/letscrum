package service

import (
	"context"
	generalv1 "github.com/letscrum/letscrum/api/general/v1"
	letscrumv1 "github.com/letscrum/letscrum/api/letscrum/v1"
	userv1 "github.com/letscrum/letscrum/api/user/v1"
	"github.com/letscrum/letscrum/internal/dao"
	"github.com/letscrum/letscrum/pkg/build"
	"github.com/letscrum/letscrum/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

type LetscrumService struct {
	letscrumv1.UnimplementedLetscrumServer
	userDao dao.UserDao
}

func NewLetscrumService(dao dao.Interface) *LetscrumService {
	return &LetscrumService{userDao: dao.UserDao()}
}

func (s *LetscrumService) GetVersion(context.Context, *emptypb.Empty) (*generalv1.GetVersionResponse, error) {
	ver := build.Version()
	return &generalv1.GetVersionResponse{
		Version: &generalv1.Version{
			Version:   ver.Version,
			GitCommit: ver.GitCommit,
			BuildDate: ver.BuildDate,
			GoVersion: ver.GoVersion,
		},
	}, nil
}

func (s *LetscrumService) SignIn(ctx context.Context, req *userv1.SignInRequest) (*userv1.SignInResponse, error) {
	user, err := s.userDao.SignIn(req.Name, req.Password)
	if err != nil {
		result := status.Convert(err)
		if result.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "get book err: %s not found", req.Name)
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}
	if user.ID == 0 {
		return nil, status.Error(codes.NotFound, "user not fount or password not correct.")
	}
	accessToken, refreshToken, errGenTokens := utils.GenerateTokens(strconv.FormatInt(user.ID, 10), user.IsSuperAdmin)
	if errGenTokens != nil {
		return nil, errGenTokens
	}
	return &userv1.SignInResponse{
		Item: &userv1.User{
			Id:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			IsSuperAdmin: user.IsSuperAdmin,
			CreatedAt:    user.CreatedAt.Unix(),
			UpdatedAt:    user.UpdatedAt.Unix(),
			Token: &userv1.Token{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		},
	}, nil
}
