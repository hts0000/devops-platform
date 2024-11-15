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

protoc.exe --proto_path %ROOT_PATH%/proto/ --plugin=protoc-gen-ts_proto=%FRONTEND_GEN_PATH%/node_modules/.bin/protoc-gen-ts_proto.cmd --ts_proto_opt=esModuleInterop=true,forceLong=long,outputEncodeMethods=false,outputPartialMethods=false,outputClientImpl=false,outputServices=false --ts_proto_out=%FRONTEND_GEN_PATH%/src/lib/%SERVICE%/api/gen/v1 --proto_path proto/%SERVICE% %SERVICE%.proto
