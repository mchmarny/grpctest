RELEASE=0.1.2

protos:
	protoc \
		--proto_path=api/v1 \
		--proto_path=$(GOOGLEAPIS_DIR) \
		--go_out=plugins=grpc:pkg/api/v1 \
		message.proto

	protoc \
		--proto_path=api/v1 \
		--proto_path=$(GOOGLEAPIS_DIR) \
		--grpc-gateway_out=logtostderr=true:pkg/api/v1 \
		message.proto

	protoc \
		--proto_path=api/v1 \
		--proto_path=$(GOOGLEAPIS_DIR) \
		--swagger_out=logtostderr=true:api/swagger/v1 \
		message.proto

run:
	go run cmd/server/main.go

call:
	go run cmd/client/main.go \
		--server=grpctest.demo.knative.tech:8080 \
		--name=John

	# curl http://localhost:8080/v1/say/john

mod:
	go mod tidy
	go mod vendor

image: mod
	gcloud builds submit \
		--project cloudylabs-public \
		--tag gcr.io/cloudylabs-public/grpctest:$(RELEASE)

.PHONY: deploy
deploy:
	kubectl apply -f deploy/service.yaml -n demo


