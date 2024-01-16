package bootstrap

import (
	"awesomeProject/internal/core-module/config"
	"awesomeProject/internal/core-module/handler"
	"awesomeProject/internal/core-module/repository"
	"awesomeProject/internal/core-module/service"
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

func Run(ctx context.Context, baseConfig *Config) error {
	log.Info("Read localConfig")
	configReader := config.NewBaseLocalConfigReader()
	localConfig, err := configReader.NewLocalConfig(baseConfig.ConfigPath)
	if err != nil {
		return err
	}

	// Read configStorage (by etcd addr)

	log.Info("Creating tarantool repository")
	bandRepository, err := repository.NewFakeBandRepository(ctx)
	if err != nil {
		return err
	}

	log.Info("Creating band service")
	bandService, err := service.NewBandService(bandRepository)
	if err != nil {
		return err
	}

	messageService := handler.NewMessageService(bandService)

	log.Info("Try to run grpc server")
	if err := runGrpcServer(localConfig.GrpcAddr, messageService); err != nil {
		return err
	}
	return nil
}

func runGrpcServer(etcdAddr string, messageService message.MessageServiceServer) error {
	listener, err := net.Listen("tcp", etcdAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	message.RegisterMessageServiceServer(grpcServer, messageService)

	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve by grpcServer: %w", err)
	}
	return nil
}
