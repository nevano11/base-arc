package bootstrap

import (
	"awesomeProject/internal/base-module/config"
	"awesomeProject/internal/base-module/handler"
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func Run(ctx context.Context, baseConfig *Config) error {
	log.Info("Read localConfig")
	configReader := config.NewBaseLocalConfigReader()
	localConfig, err := configReader.NewLocalConfig(baseConfig.ConfigPath)
	if err != nil {
		return err
	}

	// Read configStorage (by etcd addr)

	log.Info("Try to run client and send message")
	grpcConnection, err := createGrpcConnection(ctx, localConfig.GrpcAddr)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(grpcConnection)

	if err != nil {
		return err
	}

	client := message.NewMessageServiceClient(grpcConnection)

	err = handler.SendMessage(ctx, client)
	if err != nil {
		log.Errorf("Error on work of %s: %s", "SendMessage", err)
	}
	time.Sleep(2 * time.Second)

	err = handler.SumOfNumbers(ctx, client)
	if err != nil {
		log.Errorf("Error on work of %s: %s", "SumOfNumbers", err)
	}
	time.Sleep(2 * time.Second)

	err = handler.Factorial(ctx, client)
	if err != nil {
		log.Errorf("Error on work of %s: %s", "Factorial", err)
	}
	time.Sleep(2 * time.Second)

	err = handler.XPow2Chat(ctx, client)
	if err != nil {
		log.Errorf("Error on work of %s: %s", "XPow2Chat", err)
	}
	return nil
}

func createGrpcConnection(ctx context.Context, grpcAddr string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.DialContext(ctx, grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %s", err)
	}
	return conn, nil
}
