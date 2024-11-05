package main

import (
	"log"
	managerpb "manager/api/gen/v1"
	"manager/dao"
	"manager/manager"
	"shared/server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	dsn := "root:newpass@HTS.2008@tcp(127.0.0.1:3306)/devops-platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("open database connction failed", zap.Error(err))
	}

	logger.Fatal("run grpc server failed", zap.Error(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "manager",
		Addr:   ":18081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			managerpb.RegisterBizLineServiceServer(s, &manager.BizLineService{
				Logger: logger,
				MySQL:  dao.NewMySQL(db),
			})
			managerpb.RegisterProjectServiceServer(s, &manager.ProjectService{
				Logger: logger,
				MySQL:  dao.NewMySQL(db),
			})
			managerpb.RegisterAppServiceServer(s, &manager.AppService{
				Logger: logger,
				MySQL:  dao.NewMySQL(db),
			})
		},
	})))
}
