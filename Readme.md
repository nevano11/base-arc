    cmd - папка для main.go и локальных конфигов
    deployment - конфиги для docker-compose
    internal - основной код приложений
    proto - протобаф для gRPS
### План по реализации
1. [x] Научиться заворачивать в docker
2. [x] Генерация с protobuf
3. [x] Взаимодествие по grpc
4. [x] Tarantool
5. [x] Выполнение lua-скриптов
6. [ ] Развертывание в docker
7. [ ] Использование tarantool в коде
### Команды всякие
Protoc generation

    go get -u google.golang.org/protobuf/cmd/protoc-gen-go
    go install google.golang.org/protobuf/cmd/protoc-gen-go

    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

    protoc -I=proto --go_out=. proto/*.proto
    protoc -I=proto --go-grpc_out=. proto/*.proto
Tarantool tests

    docker run --name mytarantool -p3301:3301 -e TARANTOOL_USER_NAME=name -e TARANTOOL_USER_PASSWORD=pass -d -v ./init/tt:/opt/tarantool tarantool/tarantool tarantool /opt/tarantool/init.lua
    docker exec -i -t mytarantool console