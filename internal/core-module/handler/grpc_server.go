package handler

import (
	"awesomeProject/internal/core-module/service"
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"math/rand"
	"time"
)

type MessageService struct {
	bandService *service.BandService
	message.UnimplementedMessageServiceServer
}

func NewMessageService(service *service.BandService) *MessageService {
	return &MessageService{
		bandService: service,
	}
}

func (s *MessageService) SendMessage(ctx context.Context, messageFromClient *message.Message) (*message.Message, error) {
	log.Printf("Received message from client: num=%d, log=%s", messageFromClient.Num, messageFromClient.Log)

	randomBandId := rand.Intn(3) + 1
	band, err := s.bandService.GetBandById(randomBandId)
	if err != nil {
		return nil, err
	}

	return &message.Message{
		Log: band.String(),
		Num: int32(randomBandId),
	}, nil
}

func (s *MessageService) SumOfNumbers(stream message.MessageService_SumOfNumbersServer) error {
	log.Info("Start server method sum of numbers. total=0")
	total := int32(0)
	for {
		messageFromClient, err := stream.Recv()

		// stop signal
		if err == io.EOF {
			log.Infof("Received EOF signal. Result sum=%d", total)
			err := stream.SendAndClose(&message.Message{
				Log: "Server get EOF. Send result",
				Num: total,
			})

			if err != nil {
				return err
			}
			return nil
		}

		// handle err
		if err != nil {
			return err
		}

		total += messageFromClient.Num
		log.Infof("Received %d, new total=%d", messageFromClient.Num, total)
	}
}

func (s *MessageService) Factorial(messageFromClient *message.Message, stream message.MessageService_FactorialServer) error {
	num := messageFromClient.Num
	log.Infof("Start server method factorial. Client send %d", num)

	for i := num; i > 0; i-- {
		log.Infof("Server send to client %d", i)
		err := stream.Send(&message.Message{
			Log: fmt.Sprintf("Server get from client %d", num),
			Num: i,
		})
		time.Sleep(time.Second)
		if err != nil {
			return err
		}
	}

	log.Infof("Server stop streaming process")
	return nil
}

func (s *MessageService) XPow2Chat(stream message.MessageService_XPow2ChatServer) error {
	log.Infof("Start server method XPow2Chat")

	//server will wait eof
	for {
		messageFromClient, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Received EOF signal")
			return nil
		}
		if err != nil {
			return err
		}

		log.Infof("Server get x=%d", messageFromClient.Num)
		messageToSend := message.Message{
			Log: "I send to you x^2",
			Num: messageFromClient.Num * messageFromClient.Num,
		}
		err = stream.Send(&messageToSend)
		log.Infof("Server send x^2=%d", messageToSend.Num)
		if err != nil {
			return err
		}
	}
}
