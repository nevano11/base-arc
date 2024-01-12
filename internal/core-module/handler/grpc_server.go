package handler

import (
	"awesomeProject/proto/message"
	"context"
	"github.com/labstack/gommon/log"
)

type MessageService struct {
	message.UnimplementedMessageServiceServer
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) SendMessage(ctx context.Context, messageFromClient *message.Message) (*message.Message, error) {
	log.Printf("Received message from client: num=%d, log=%s", messageFromClient.Num, messageFromClient.Log)
	return &message.Message{
		Log: "I receive your message. Send x2",
		Num: messageFromClient.Num * 2,
	}, nil
}
