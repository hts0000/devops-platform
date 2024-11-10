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


# Path to this cli protobufjs-cli
PROTOC_GEN_JS_PATH="${FRONTEND_GEN_PATH}/node_modules/.bin/pbjs"
PROTOC_GEN_TS_PATH="${FRONTEND_GEN_PATH}/node_modules/.bin/pbts"

gen_ts_pb() {
       local service_name="$1"

       if [[ ! -d "${FRONTEND_GEN_PATH}/api/gen/v1/${service_name}" ]]; then
              mkdir -p "${FRONTEND_GEN_PATH}/api/gen/v1/${service_name}"
       fi
       
       ${PROTOC_GEN_JS_PATH} -t static-module -w es6 ${ROOT_PATH}/proto/${service_name}/${service_name}.proto \
              --no-create --no-encode --no-verify --no-delimited \
              -o ${FRONTEND_GEN_PATH}/api/gen/v1/${service_name}/${service_name}_pb.js

       ${PROTOC_GEN_TS_PATH} ${FRONTEND_GEN_PATH}/api/gen/v1/${service_name}/${service_name}_pb.js \
              -o ${FRONTEND_GEN_PATH}/api/gen/v1/${service_name}/${service_name}_pb.d.ts

       # Path to this plugin ts-protoc-gen
       # PROTOC_GEN_TS_PATH="${FRONTEND_GEN_PATH}/node_modules/.bin/protoc-gen-ts"

       # protoc --proto_path ${ROOT_PATH}/proto/ \
       #        --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
       #        --js_out="import_style=commonjs,binary:${FRONTEND_GEN_PATH}/api/gen/v1/hello" \
       #        --ts_out="${FRONTEND_GEN_PATH}/api/gen/v1/hello" \
       #        --proto_path proto/hello \
       #        proto/hello/hello.proto
}

gen() {
       local services="$@"
       for service in ${services}; do
              gen_go_pb ${service}
              # gen_ts_pb ${service}
       done
}

gen "hello" "manager"