package bootstrap

import (
	"awesomeProject/internal/base-module/config"
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	err = sendMessage(client)
	if err != nil {
		return err
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

func sendMessage(client message.MessageServiceClient) error {
	response, err := client.SendMessage(context.Background(), &message.Message{Log: "Hello From Client!", Num: 21})
	if err != nil {
		return fmt.Errorf("error when calling SendMessage: %s", err)
	}
	fmt.Printf("Response from server: num=%d, log=%s", response.Num, response.Log)
	return nil
}
