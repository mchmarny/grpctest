


protoc \
    --proto_path=${GOOGLEAPIS_DIR} \
    --proto_path=. \
    helloworld/helloworld.proto \
    --go_out=plugins=grpc:.

