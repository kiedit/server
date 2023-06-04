package queue

import (
	"kiedit/media"
	"log"

	"github.com/streadway/amqp"
)

func (self *QueueStruct) Publish(splitVideoInput media.SplitVideoInput) error {
	err := self.Channel.Publish(
		"",       // exchange
		"upload", // key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(splitVideoInput.OutputDirPath),
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
