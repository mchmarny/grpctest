
# proto
protoc \
    --proto_path=api/v1 \
    --proto_path=${GOOGLEAPIS_DIR} \
    --go_out=plugins=grpc:pkg/api/v1 \
    message.proto

protoc \
    --proto_path=api/v1 \
    --proto_path=${GOOGLEAPIS_DIR} \
    --grpc-gateway_out=logtostderr=true:pkg/api/v1 \
    message.proto

protoc \
    --proto_path=api/v1 \
    --proto_path=${GOOGLEAPIS_DIR} \
    --swagger_out=logtostderr=true:api/swagger/v1 \
    message.proto





# server
go run cmd/server/grpc/main.go
go run cmd/server/http/main.go

# client
go run cmd/client/main.go

# rest test
curl http://localhost:7777/v1/hello/hithere