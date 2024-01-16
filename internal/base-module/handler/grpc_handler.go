package handler

import (
	"awesomeProject/proto/message"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"time"
)

func SendMessage(ctx context.Context, client message.MessageServiceClient) error {
	messageToSend := &message.Message{Log: "Hello From Client!", Num: int32(rand.Intn(100) + 1)}

	log.Infof("Client send message to server: %s", messageToSend)
	response, err := client.SendMessage(ctx, messageToSend)
	if err != nil {
		log.Errorf("error when calling SendMessage: %s", err)
		return fmt.Errorf("error when calling SendMessage: %s", err)
	}
	log.Infof("Response from server: num=%d, log=%s", response.Num, response.Log)
	return nil
}

func SumOfNumbers(ctx context.Context, client message.MessageServiceClient) error {
	log.Infof("Client run SumOfNumbers")
	numberStream, err := client.SumOfNumbers(ctx, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	for i := int32(1); i < int32(5); i++ {
		log.Infof("Client send to server %d", i)
		err := numberStream.Send(&message.Message{
			Log: fmt.Sprintf("I send to server %d", i),
			Num: i,
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

func Factorial(ctx context.Context, client message.MessageServiceClient) error {
	messageForServer := message.Message{
		Log: "I want to get 4 3 2 1",
		Num: 5,
	}
	log.Infof("Client run Factorial with message: %d, %s", messageForServer.Num, messageForServer.Log)

	factorialStream, err := client.Factorial(ctx, &messageForServer, grpc.EmptyCallOption{})
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

func XPow2Chat(ctx context.Context, client message.MessageServiceClient) error {
	log.Infof("Client run XPow2Chat")

	chatStream, err := client.XPow2Chat(ctx, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	// Client want to get x^2 for numbers 2, 4, 5, -10, 2, 11, 0, 6, 7
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
	numbers := []int32{2, 4, 5, -10, 2, 11, 0, 6, 7}
	for _, v := range numbers {
		messageToSend := message.Message{
			Log: "I want to find x^2",
			Num: v,
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
