package queue

import (
	"log"

	"github.com/streadway/amqp"
)

type QueueStruct struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func (self *QueueStruct) Connect() error {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
