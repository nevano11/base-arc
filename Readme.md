    cmd - папка для main.go и локальных конфигов
    deployment - конфиги для docker-compose
    internal - основной код приложений
    proto - протобаф для gRPS
### План по реализации
1. [x] Научиться заворачивать в docker
2. [x] Генерация с protobuf
3. [ ] Взаимодествие по grpc
4. [ ] Tarantool
5. [ ] Выполнение lua-скриптов
6. [ ] Развертывание в docker
### Команды всякие
   protoc -I=proto --go_out=. proto/*.proto