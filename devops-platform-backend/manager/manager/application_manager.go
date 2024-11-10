package manager

import (
	"context"
	managerpb "manager/api/gen/v1"
	"manager/dao"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppService struct {
	MySQL  *dao.MySQL
	Logger *zap.Logger
	managerpb.UnimplementedAppServiceServer
}

func (s *AppService) CreateProject(context.Context, *managerpb.CreateAppRequest) (*managerpb.CreateAppResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *AppService) DeleteProject(context.Context, *managerpb.DeleteAppRequest) (*managerpb.DeleteAppResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *AppService) UpdateProject(context.Context, *managerpb.UpdateAppRequest) (*managerpb.UpdateAppResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *AppService) GetProject(context.Context, *managerpb.GetAppRequest) (*managerpb.GetAppResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *AppService) GetProjects(context.Context, *managerpb.GetAppsRequest) (*managerpb.GetAppsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
