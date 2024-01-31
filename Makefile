all: build

build:
	go build 

proto:
	protoc --go_out=./pb --go-grpc_out=./pb proto/*.proto