# grpctest


## generate

### proto

```shell
protoc \
    --proto_path=api/v1 \
    --proto_path=${GOOGLEAPIS_DIR} \
    --go_out=plugins=grpc:pkg/api/v1 \
    message.proto
```

### gateway

```shell
protoc \
    --proto_path=api/v1 \
    --proto_path=${GOOGLEAPIS_DIR} \
    --grpc-gateway_out=logtostderr=true:pkg/api/v1 \
    message.proto
```

### docs

```shell
protoc \
    --proto_path=api/v1 \
    --proto_path=${GOOGLEAPIS_DIR} \
    --swagger_out=logtostderr=true:api/swagger/v1 \
    message.proto
```


## run

### locally

```shell
go run cmd/server/main.go
```

### cloud run on gke

```shell
kubectl apply -f deploy/service.yaml -n demo
```

## client

### local (grpc)

```shell
go run cmd/client/main.go --server=":8080" --name=John
```

### local (rest)

```shell
curl http://localhost:7777/v1/say/mark
```

### cloud run on gke

```shell
go run cmd/client/main.go \
    --server=grpctest.demo.knative.tech:8080 \
    --name=John
```