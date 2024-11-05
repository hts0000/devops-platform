@echo off

set ROOT_PATH=%cd%
set FRONTEND_GEN_PATH=%ROOT_PATH%\devops-platform-frontend
set BACKEND_GEN_PATH=%ROOT_PATH%\devops-platform-backend

set SERVICE=hello
protoc.exe --proto_path %ROOT_PATH%/proto/ --go_out %BACKEND_GEN_PATH%/%SERVICE%/api/gen/v1 --go_opt paths=source_relative --go-grpc_out %BACKEND_GEN_PATH%/%SERVICE%/api/gen/v1 --go-grpc_opt paths=source_relative --proto_path proto/%SERVICE% %SERVICE%.proto

protoc.exe --proto_path %ROOT_PATH%/proto/ --grpc-gateway_out %BACKEND_GEN_PATH%/%SERVICE%/api/gen/v1 --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --proto_path proto/%SERVICE% %SERVICE%.proto

set SERVICE=manager
protoc.exe --proto_path %ROOT_PATH%/proto/ --go_out %BACKEND_GEN_PATH%/%SERVICE%/api/gen/v1 --go_opt paths=source_relative --go-grpc_out %BACKEND_GEN_PATH%/%SERVICE%/api/gen/v1 --go-grpc_opt paths=source_relative --proto_path proto/%SERVICE% %SERVICE%.proto

protoc.exe --proto_path %ROOT_PATH%/proto/ --grpc-gateway_out %BACKEND_GEN_PATH%/%SERVICE%/api/gen/v1 --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --proto_path proto/%SERVICE% %SERVICE%.proto
