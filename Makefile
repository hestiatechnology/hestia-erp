all: run

run:
	go run main.go

build:
	go build ./...

test:
	go test ./...

proto:
	protoc --go_out=. --go_opt=module=hestia/api --go-grpc_out=. --go-grpc_opt=module=hestia/api  proto/*.proto

clean:
	git clean -fdx