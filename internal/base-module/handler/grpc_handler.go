package handler

import (
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"io"
	"time"
)

type Handler struct {
	client message.MessageServiceClient
}

func NewHandler(client message.MessageServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) SendMessage(ctx context.Context, messageToSend *message.Message) error {
	log.Infof("Client send message to server: %s", messageToSend)
	response, err := h.client.SendMessage(ctx, messageToSend)
	if err != nil {
		log.Errorf("error when calling SendMessage: %s", err)
		return fmt.Errorf("error when calling SendMessage: %s", err)
	}
	log.Infof("Response from server: num=%d, log=%s", response.Num, response.Log)
	return nil
}

func (h *Handler) SumOfNumbers(ctx context.Context, numbers []int) error {
	log.Infof("Client run SumOfNumbers")
	numberStream, err := h.client.SumOfNumbers(ctx, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	for _, num := range numbers {
		log.Infof("Client send to server %d", num)
		err := numberStream.Send(&message.Message{
			Log: fmt.Sprintf("I send to server %d", num),
			Num: int32(num),
		})
		time.Sleep(time.Second)
		if err != nil {
			return err
		}
	}
	log.Info("Client try to get answer from server")
	result, err := numberStream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Infof("Result from server: %s", result)
	return nil
}

func (h *Handler) Factorial(ctx context.Context, messageToSend *message.Message) error {
	log.Infof("Client run Factorial with message: %d, %s", messageToSend.Num, messageToSend.Log)

	factorialStream, err := h.client.Factorial(ctx, messageToSend, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	for {
		messageFromServer, err := factorialStream.Recv()

		if err == io.EOF {
			log.Infof("Client get EOF")
			return nil
		}
		if err != nil {
			return err
		}

		log.Infof("Client get from server message: Num=%d Log=%s", messageFromServer.Num, messageFromServer.Log)
	}

	// return on for cycle
}

func (h *Handler) XPow2Chat(ctx context.Context, numbers []int) error {
	log.Infof("Client run XPow2Chat")

	chatStream, err := h.client.XPow2Chat(ctx, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	// Client want to get x^2 for numbers
	// reading
	go func() {
		for {
			messageFromServer, err := chatStream.Recv()
			if err == io.EOF {
				log.Infof("Client get EOF")
				return
			}
			if err != nil {
				return
			}
			log.Infof("Client get from server message: Num=%d Log=%s", messageFromServer.Num, messageFromServer.Log)
		}
	}()

	// sending
	for _, num := range numbers {
		messageToSend := message.Message{
			Log: "I want to find x^2",
			Num: int32(num),
		}
		log.Infof("Client send to server message: Num=%d Log=%s", messageToSend.Num, messageToSend.Log)
		err := chatStream.Send(&messageToSend)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	// stop sending
	err = chatStream.CloseSend()
	if err != nil {
		return err
	}
	return nil
}
