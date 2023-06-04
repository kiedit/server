package queue

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

type QueueStruct struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func (self *QueueStruct) Connect() error {
	var RABBITMQ_URL_HOST_FROM_ENV = os.Getenv("RABBITMQ_URL_HOST")

	var RABBITMQ_URL_HOST = "localhost:5672/"

	if len(RABBITMQ_URL_HOST_FROM_ENV) > 0 {
		RABBITMQ_URL_HOST = RABBITMQ_URL_HOST_FROM_ENV
	}

	var RABBITMQ_URL = "amqp://guest:guest@" + RABBITMQ_URL_HOST

	connection, err := amqp.Dial(RABBITMQ_URL)
	if err != nil {
		log.Println("Connection Dial error")
		return err
	}
	self.Connection = connection

	channel, err := connection.Channel()
	if err != nil {
		log.Println("Connection Channel error")
		return err
	}
	self.Channel = channel

	queue, err := channel.QueueDeclare(
		"upload",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Channel Queue Declare error")
		return err
	}
	self.Queue = queue

	return nil
}

func (self *QueueStruct) Close() error {
	if err := self.Connection.Close(); err != nil {
		return err
	}
	if err := self.Channel.Close(); err != nil {
		return err
	}
	return nil
}
