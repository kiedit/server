package queue

import (
	"log"

	"github.com/streadway/amqp"
)

func (self *QueueStruct) Publish(data string) error {
	err := self.Channel.Publish(
		"",       // exchange
		"upload", // key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)
	if err != nil {
		log.Println("Channel Publish error")
		return err
	}

	log.Println("Queue status:", self.Queue)
	log.Println("Successfully published message")
	return nil
}
