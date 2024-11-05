package manager

import (
	"context"
	managerpb "manager/api/gen/v1"
	"manager/dao"

	"go.uber.org/zap"
)

type AppService struct {
	MySQL  *dao.MySQL
	Logger *zap.Logger
	managerpb.UnimplementedAppServiceServer
}

func (s *AppService) CreateProject(context.Context, *managerpb.CreateAppRequest) (*managerpb.CreateAppResponse, error)
func (s *AppService) DeleteProject(context.Context, *managerpb.DeleteAppRequest) (*managerpb.DeleteAppResponse, error)
func (s *AppService) UpdateProject(context.Context, *managerpb.UpdateAppRequest) (*managerpb.UpdateAppResponse, error)
func (s *AppService) GetProject(context.Context, *managerpb.GetAppRequest) (*managerpb.GetAppResponse, error)
func (s *AppService) GetProjects(context.Context, *managerpb.GetAppsRequest) (*managerpb.GetAppsResponse, error)
