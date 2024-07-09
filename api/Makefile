all: run

run:
	go run main.go

build:
	go build -v ./...

test:
	go test -v ./...

proto-server:
	protoc --proto_path=..\proto\v1 --go_out=. --go_opt=module=hestia/api --go-grpc_out=. --go-grpc_opt=module=hestia/api ..\proto\v1\*.proto

clean:
	git clean -fdx