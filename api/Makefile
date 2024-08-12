all: run

run:
	go run main.go

build:
	go build -v hestia/api

build-tools:
	go build -v -o ./hst-protextil-migration ./cmd/protextil-migration
	go build -v -o ./hst-macwin-migration ./cmd/macwin-migration

build-tools-win:
	go build -v -o ./hst-protextil-migration.exe ./cmd/protextil-migration
	go build -v -o ./hst-macwin-migration.exe ./cmd/macwin-migration

test:
	go test -v ./...

proto-server:
	protoc --proto_path=../proto/v1 --go_out=. --go_opt=module=hestia/api --go-grpc_out=. --go-grpc_opt=module=hestia/api ../proto/v1/*.proto

clean:
	git clean -Xfd