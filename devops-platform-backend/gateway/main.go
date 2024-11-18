package main

import (
	"context"
	hellopb "hello/api/gen/v1"
	"io"
	"log"
	managerpb "manager/api/gen/v1"
	"net/http"
	"shared/server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"

	// 需要匿名导入errdetails，用以获取errdetails中rpc错误映射到http错误的方法
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseEnumNumbers: true, // 返回常量时，返回对应的number而不是string
					UseProtoNames:  true, // 返回JSON时，字段名格式为xxx_xxx格式
				},
			},
		),
		runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
			// TODO: handle special request
			return runtime.DefaultHeaderMatcher(s)
		}),
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := metadata.MD{}
			// TODO: add token from cookie to grpc metadata
			// token, err := r.Cookie("token")
			// if err != nil {
			// 	return md
			// }
			// md.Set("token", token.Value)
			return md
		}),
		runtime.WithMiddlewares(func(hf runtime.HandlerFunc) runtime.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
				// TODO: add traceid
				b, err := io.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				body := string(b)
				logger.Info("received request", zap.String("host", r.URL.Host), zap.String("scheme", r.URL.Scheme), zap.String("path", r.URL.Path), zap.String("body", body), zap.String("traceid", ""))
				hf(w, r, pathParams)
				logger.Info("handled request", zap.String("host", r.URL.Host), zap.String("scheme", r.URL.Scheme), zap.String("path", r.URL.Path), zap.String("traceid", ""))
			}
		}),
	)

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "hello",
			addr:         "localhost:28081",
			registerFunc: hellopb.RegisterGreeterHandlerFromEndpoint,
		},
		{
			name:         "bizline",
			addr:         "localhost:18081",
			registerFunc: managerpb.RegisterBizLineServiceHandlerFromEndpoint,
		},
		// {
		// 	name:         "project",
		// 	addr:         "localhost:18081",
		// 	registerFunc: managerpb.RegisterBizLineServiceHandlerFromEndpoint,
		// },
		// {
		// 	name:         "application",
		// 	addr:         "localhost:18081",
		// 	registerFunc: managerpb.RegisterBizLineServiceHandlerFromEndpoint,
		// },
	}

	for _, s := range serverConfig {
		err := s.registerFunc(
			c, mux, s.addr,
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)
		if err != nil {
			logger.Fatal("cannot register service", zap.String("service", s.name), zap.Error(err))
		}
	}

	addr := ":18080"
	logger.Info("grpc gateway started", zap.String("addr", addr))
	logger.Fatal("cannot listen and server", zap.Error(http.ListenAndServe(addr, cors.Default().Handler(mux))))
}
