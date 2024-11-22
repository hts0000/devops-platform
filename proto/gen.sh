ROOT_PATH=$(cd $(dirname $0)/../ && pwd)
FRONTEND_GEN_PATH="${ROOT_PATH}/devops-platform-frontend"
BACKEND_GEN_PATH="${ROOT_PATH}/devops-platform-backend"

gen_go_pb() {
       local service_name="$1"

       if [[ ! -d "${BACKEND_GEN_PATH}/${service_name}/api/gen/v1" ]]; then
              mkdir -p "${BACKEND_GEN_PATH}/${service_name}/api/gen/v1"
       fi

       protoc --proto_path ${ROOT_PATH}/proto/ \
              --go_out ${BACKEND_GEN_PATH}/${service_name}/api/gen/v1 \
              --go_opt paths=source_relative \
              --go-grpc_out ${BACKEND_GEN_PATH}/${service_name}/api/gen/v1 \
              --go-grpc_opt paths=source_relative \
              --proto_path proto/${service_name} \
              ${service_name}.proto

       protoc --proto_path ${ROOT_PATH}/proto/ \
              --grpc-gateway_out ${BACKEND_GEN_PATH}/${service_name}/api/gen/v1 \
              --grpc-gateway_opt logtostderr=true \
              --grpc-gateway_opt paths=source_relative \
              --proto_path proto/${service_name} \
              ${service_name}.proto
}

gen_ts_pb() {
       local service_name="$1"
       local gen_path="${FRONTEND_GEN_PATH}/src/lib/${service_name}/api/gen/v1"
       local ts_gen_path="${FRONTEND_GEN_PATH}/node_modules/.bin/protoc-gen-ts_proto"
       
       if [[ ! -f "${ts_gen_path}" ]]; then
              echo "cant not find <${ts_gen_path}>, gen ts file failed"
              exit 1
       fi

       if [[ ! -d "${gen_path}" ]]; then
              mkdir -p "${gen_path}"
       fi

       protoc --proto_path ${ROOT_PATH}/proto/ \
              --plugin=protoc-gen-ts_proto=${ts_gen_path} \
              --ts_proto_opt=esModuleInterop=true,forceLong=string,outputEncodeMethods=false,outputPartialMethods=false,outputClientImpl=false,outputServices=false \
              --ts_proto_out=${gen_path} \
              --proto_path proto/${service_name} \
              ${service_name}.proto

       # remove useless validate and google api dir
       # frontend project dont need this
       rm -rf ${gen_path}/buf ${gen_path}/google &> /dev/null
}

gen() {
       local services="$@"
       for service in ${services}; do
              gen_go_pb ${service}
              gen_ts_pb ${service}
       done
}

gen "hello" "manager"