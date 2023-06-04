package queue

import (
	"encoding/json"
	"kiedit/media"
	"log"

	"github.com/streadway/amqp"
)

func (self *QueueStruct) Publish(splitVideoInput media.SplitVideoInput) error {
	var body, err = json.Marshal(splitVideoInput)

	if err != nil {
		log.Println("Channel Publish Data error")
		return err
	}

	err = self.Channel.Publish(
		"",       // exchange
		"upload", // key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Channel Publish error")
		return err
	}

	log.Println("Successfully published message")
	return nil
}
