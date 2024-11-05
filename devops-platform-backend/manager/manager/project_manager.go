package manager

import (
	"context"
	managerpb "manager/api/gen/v1"
	"manager/dao"

	"go.uber.org/zap"
)

type ProjectService struct {
	MySQL  *dao.MySQL
	Logger *zap.Logger
	managerpb.UnimplementedProjectServiceServer
}

func (s *ProjectService) CreateProject(context.Context, *managerpb.CreateProjectRequest) (*managerpb.CreateProjectResponse, error)
func (s *ProjectService) DeleteProject(context.Context, *managerpb.DeleteProjectRequest) (*managerpb.DeleteProjectResponse, error)
func (s *ProjectService) UpdateProject(context.Context, *managerpb.UpdateProjectRequest) (*managerpb.UpdateProjectResponse, error)
func (s *ProjectService) GetProject(context.Context, *managerpb.GetProjectRequest) (*managerpb.GetProjectResponse, error)
func (s *ProjectService) GetProjects(context.Context, *managerpb.GetProjectsRequest) (*managerpb.GetProjectsResponse, error)
