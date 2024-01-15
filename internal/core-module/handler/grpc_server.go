package handler

import (
	"awesomeProject/internal/core-module/service"
	"awesomeProject/proto/message"
	"context"
	"github.com/labstack/gommon/log"
	"math/rand"
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
