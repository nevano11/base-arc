package bootstrap

import (
	local_config "awesomeProject/internal/base-module/config"
	"awesomeProject/internal/base-module/handler"
	"awesomeProject/pkg/config"
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"time"
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
	err := readLocalConfig(pathConfig.ConfigPath, pathConfig.ConfigFilename)
	if err != nil {
		return err
	}

	// Read configStorage (by etcd addr)

	log.Info("Try to run client and send message")
	grpcConnection, err := createGrpcConnection(ctx, AppConfig.GrpcAddr)
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
	grpcHandler := handler.NewHandler(client)

	err = grpcHandler.SendMessage(ctx, &message.Message{Log: "Hello From Client!", Num: int32(rand.Intn(100) + 1)})
	if err != nil {
		log.Errorf("Error on work of %s: %s", "SendMessage", err)
	}
	time.Sleep(2 * time.Second)

	err = grpcHandler.SumOfNumbers(ctx, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if err != nil {
		log.Errorf("Error on work of %s: %s", "SumOfNumbers", err)
	}
	time.Sleep(2 * time.Second)

	err = grpcHandler.Factorial(ctx, &message.Message{
		Log: "I want to get 5 4 3 2 1",
		Num: 5,
	})
	if err != nil {
		log.Errorf("Error on work of %s: %s", "Factorial", err)
	}
	time.Sleep(2 * time.Second)

	err = grpcHandler.XPow2Chat(ctx, []int{1, 2, -10, 5, 75, 11, -100, 2})
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
