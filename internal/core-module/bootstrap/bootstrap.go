package bootstrap

import (
	"awesomeProject/internal/core-module/config"
	"awesomeProject/internal/core-module/handler"
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

	log.Info("Try to run grpc server")
	if err := runGrpcServer(localConfig.GrpcAddr); err != nil {
		return err
	}
	return nil
}

func runGrpcServer(etcdAddr string) error {
	listener, err := net.Listen("tcp", etcdAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	messageService := handler.NewMessageService()
	message.RegisterMessageServiceServer(grpcServer, messageService)

	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve by grpcServer: %w", err)
	}
	return nil
}
