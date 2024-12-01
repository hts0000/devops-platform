package manager

import (
	"context"
	"errors"
	managerpb "manager/api/gen/v1"
	"manager/dao"
	"shared/model"

	"github.com/bufbuild/protovalidate-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

	biz, err := s.MySQL.GetBizlineByName(ctx, req.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Logger.Error("get bizline by name failed", zap.String("name", req.Name), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	if biz != nil && biz.ID != 0 {
		s.Logger.Error("bizline already exist", zap.String("name", req.Name))
		return nil, status.Error(codes.AlreadyExists, "bizline already exist")
	}

	biz = &model.BizLine{
		Name:          req.Name,
		ResponsibleID: uint(req.ResponsibleId),
		Description:   req.Description,
	}
	if err := s.MySQL.CreateBizLine(ctx, biz); err != nil {
		s.Logger.Error("create bizline failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	res := &managerpb.CreateBizLineResponse{
		Id: uint64(biz.ID),
	}
	return res, nil
}

func (s *BizLineService) DeleteBizLine(ctx context.Context, req *managerpb.DeleteBizLineRequest) (*managerpb.DeleteBizLineResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("delete bizline req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	biz, err := s.MySQL.GetBizLineByID(ctx, uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Logger.Error("delete bizline failed, record not found", zap.Uint("id", uint(req.Id)))
			return nil, status.Error(codes.NotFound, "record not found")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	_, err = s.MySQL.DeleteBizLine(ctx, biz.ID)
	if err != nil {
		s.Logger.Error("delete bizline failed", zap.Uint("id", uint(req.Id)), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	s.Logger.Info("delete bizline success", zap.Any("bizline", biz))

	res := &managerpb.DeleteBizLineResponse{
		Id: uint64(biz.ID),
	}
	return res, nil
}

func (s *BizLineService) UpdateBizLine(ctx context.Context, req *managerpb.BizLineEntity) (*managerpb.UpdateBizLineResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("update bizline req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	biz, err := s.MySQL.GetBizLineByID(ctx, uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Logger.Warn("bizline not found", zap.Uint("id", uint(req.Id)))
			return nil, status.Error(codes.NotFound, "bizline not found")
		}

		s.Logger.Error("get bizline by id failed", zap.Uint("id", uint(req.Id)), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	biz.Name = req.BizLine.Name
	biz.ResponsibleID = uint(req.BizLine.ResponsibleId)
	biz.Description = req.BizLine.Description
	if err := s.MySQL.UpdateBizLine(ctx, biz); err != nil {
		s.Logger.Error("update bizline failed", zap.Any("req", req), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	s.Logger.Info("update bizline success", zap.Any("bizline", biz))

	res := &managerpb.UpdateBizLineResponse{}
	return res, nil
}

func (s *BizLineService) GetBizLine(ctx context.Context, req *managerpb.GetBizLineRequest) (*managerpb.BizLineEntity, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("get bizline req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	biz, err := s.MySQL.GetBizLineByID(ctx, uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Logger.Warn("bizline not found", zap.Uint("id", uint(req.Id)))
			return nil, status.Error(codes.NotFound, "bizline not found")
		}

		s.Logger.Error("get bizline by id failed", zap.Uint("id", uint(req.Id)), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	s.Logger.Info("get bizline by id success", zap.Any("bizline", biz))

	res := &managerpb.BizLineEntity{
		Id: uint64(biz.ID),
		BizLine: &managerpb.BizLine{
			Name:          biz.Name,
			ResponsibleId: uint64(biz.ResponsibleID),
			Description:   biz.Description,
			CreatedAt:     int32(biz.CreatedAt.Unix()),
		},
	}
	return res, nil
}

func (s *BizLineService) GetBizLines(ctx context.Context, req *managerpb.GetBizLinesRequest) (*managerpb.GetBizLinesResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("get bizlines req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	bizs, total, err := s.MySQL.GetBizLines(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		s.Logger.Error("get bizlines failed", zap.Int64("page", req.Page), zap.Int64("pageSize", req.PageSize), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if total == 0 || len(bizs) == 0 {
		s.Logger.Warn("bizlines has not record", zap.Int64("total", total), zap.Int("bizsLen", len(bizs)))
		return nil, status.Error(codes.OK, "bizlines has not record")
	}

	res := &managerpb.GetBizLinesResponse{
		TotalCount: total,
		Bizlines:   make([]*managerpb.BizLineEntity, 0, len(bizs)),
	}

	for _, biz := range bizs {
		res.Bizlines = append(res.Bizlines, &managerpb.BizLineEntity{
			Id: uint64(biz.ID),
			BizLine: &managerpb.BizLine{
				Name:          biz.Name,
				ResponsibleId: uint64(biz.ResponsibleID),
				Description:   biz.Description,
				CreatedAt:     int32(biz.CreatedAt.Unix()),
			},
		})
	}

	return res, nil
}
