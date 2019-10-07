GOOS?=linux
SERVER_PORT?=50051
GOOGLE_APPLICATION_CREDENTIALS?=""
BINARY_OUTPUT?="bin/image-service"

protoc:
	protoc -I proto/v1 --go_out=plugins=grpc:pkg/imagepb/v1 proto/v1/service.proto

build:
	CGO_ENABLED=0 GOOS=${GOOS} go build -a -installsuffix cgo -o ${BINARY_OUTPUT} cmd/server/main.go

run:
	SERVER_PORT=${SERVER_PORT} GOOGLE_APPLICATION_CREDENTIALS=${GOOGLE_APPLICATION_CREDENTIALS} ./bin/image-service

image:
	docker build -t github.com/dimitriin/image-service:latest .