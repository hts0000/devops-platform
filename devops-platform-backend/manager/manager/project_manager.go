package manager

import (
	"context"
	managerpb "manager/api/gen/v1"
	"manager/dao"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectService struct {
	MySQL  *dao.MySQL
	Logger *zap.Logger
	managerpb.UnimplementedProjectServiceServer
}

func (s *ProjectService) CreateProject(context.Context, *managerpb.CreateProjectRequest) (*managerpb.CreateProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *ProjectService) DeleteProject(context.Context, *managerpb.DeleteProjectRequest) (*managerpb.DeleteProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *ProjectService) UpdateProject(context.Context, *managerpb.UpdateProjectRequest) (*managerpb.UpdateProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *ProjectService) GetProject(context.Context, *managerpb.GetProjectRequest) (*managerpb.GetProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *ProjectService) GetProjects(context.Context, *managerpb.GetProjectsRequest) (*managerpb.GetProjectsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
