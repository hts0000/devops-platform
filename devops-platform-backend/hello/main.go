package main

import (
	hellopb "hello/api/gen/v1"
	"hello/dao"
	"hello/hello"
	"log"
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
		Name:   "hello",
		Addr:   ":28081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			hellopb.RegisterGreeterServer(s, &hello.Service{
				MySQL:  dao.NewMySQL(db),
				Logger: logger,
			})
		},
	})))
}
