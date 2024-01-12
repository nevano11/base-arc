include deployment/.env

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