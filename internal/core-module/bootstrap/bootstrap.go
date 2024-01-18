package bootstrap

import (
	local_config "awesomeProject/internal/core-module/config"
	"awesomeProject/internal/core-module/handler"
	"awesomeProject/internal/core-module/repository"
	"awesomeProject/internal/core-module/service"
	"awesomeProject/pkg/config"
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

var AppConfig *local_config.Config

func init() {
	AppConfig = &local_config.Config{}
}

func readLocalConfig(configPath, configFilename string) error {
	log.Infof("Read localConfig by path=%s", configPath)
	configReader := config.NewConfigReader(configPath, configFilename)

	err := configReader.ReadConfig(AppConfig)
	return err
}

func Run(ctx context.Context, pathConfig *PathConfig) error {
	// Read localConfig
	err := readLocalConfig(pathConfig.ConfigPath, pathConfig.ConfigFilename)
	if err != nil {
		return err
	}

	// Read configStorage (by etcd addr)

	// Create repositories and services
	// Repositories
	log.Info("Creating band repository")
	bandRepository, err := repository.NewMockBandRepository(ctx)
	if err != nil {
		return err
	}

	// Services
	log.Info("Creating band service")
	bandService, err := service.NewBandService(bandRepository)
	if err != nil {
		return err
	}

	// Handlers
	log.Info("Creating message service handler")
	messageService := handler.NewMessageService(bandService)

	log.Info("Try to run grpc server")
	if err := runGrpcServer(AppConfig.GrpcAddr, messageService); err != nil {
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
