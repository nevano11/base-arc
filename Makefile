include deployment/.env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

run:
	make protoc/install
	make tarantool/down
	make tarantool/initialize

protoc/install:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go

	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

	protoc -I=proto --go_out=. proto/*.proto
	protoc -I=proto --go-grpc_out=. proto/*.proto

tarantool/initialize:
	docker run --name mytarantool -p3301:3301 -e TARANTOOL_USER_NAME=name -e TARANTOOL_USER_PASSWORD=pass -d -v ./init/tt:/opt/tarantool tarantool/tarantool tarantool /opt/tarantool/init.lua

tarantool/down:
	docker rm mytarantool --force