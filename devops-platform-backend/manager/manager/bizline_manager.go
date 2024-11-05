package manager

import (
	"context"
	managerpb "manager/api/gen/v1"
	"manager/dao"

	"github.com/bufbuild/protovalidate-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BizLineService struct {
	MySQL  *dao.MySQL
	Logger *zap.Logger
	managerpb.UnimplementedBizLineServiceServer
}

func (s *BizLineService) CreateBizLine(ctx context.Context, req *managerpb.BizLine) (*managerpb.CreateBizLineResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("create bizline req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
}

func (s *BizLineService) DeleteBizLine(ctx context.Context, req *managerpb.DeleteBizLineRequest) (*managerpb.DeleteBizLineResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("delete bizline req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

}

func (s *BizLineService) UpdateBizLine(ctx context.Context, req *managerpb.BizLineEntity) (*managerpb.UpdateBizLineResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("update bizline req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
}

func (s *BizLineService) GetBizLine(ctx context.Context, req *managerpb.GetBizLineRequest) (*managerpb.BizLineEntity, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("get bizline req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
}

func (s *BizLineService) GetBizLines(ctx context.Context, req *managerpb.GetBizLinesRequest) (*managerpb.GetBizLinesResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("get bizlines req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
}
