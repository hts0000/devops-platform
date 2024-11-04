package hello

import (
	"context"
	hellopb "hello/api/gen/v1"
	"hello/dao"
	"shared/model"

	"github.com/bufbuild/protovalidate-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	MySQL  *dao.MySQL
	Logger *zap.Logger
	hellopb.UnimplementedGreeterServer
}

func (s *Service) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	err := protovalidate.Validate(req)
	if err != nil {
		s.Logger.Error("req validate failed", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	s.Logger.Info("received hello request", zap.String("name", req.Name))
	s.Logger.Info("create say hello record")

	r := &model.SayHelloRecord{
		Name: req.Name,
	}
	err = s.MySQL.CreateSayHelloRecord(r)
	if err != nil {
		s.Logger.Error("create say hello record failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	s.Logger.Info("create say hello record success", zap.Uint("id", r.ID))

	err = s.MySQL.DeleteSayHelloRecord(r)
	if err != nil {
		s.Logger.Error("delete say hello record failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	s.Logger.Info("delete say hello record success", zap.Uint("id", r.ID))

	// TODO: rich error info
	return &hellopb.HelloResponse{
		Message: "Hello!",
	}, nil
}
