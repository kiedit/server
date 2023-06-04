package queue

import (
	"encoding/json"
	"kiedit/media"
	"log"
)

func (self *QueueStruct) Consume() error {
	msgs, err := self.Channel.Consume(
		"upload", // queue
		"",       // consumer
		true,     // auto ack
		false,    // exclusive
		false,    // no local
		false,    // no wait
		nil,      //args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Println("Received Message: %s\n", string(msg.Body))

			if err := processQueueTask(msg.Body); err != nil {
				log.Println("Failed media split: %s\n", string(msg.Body))
			}
		}
	}()

	log.Println("Waiting for messages...")
	<-forever

	return nil
}

func processQueueTask(body []byte) error {
	var splitVideoInput media.SplitVideoInput

	if err := json.Unmarshal(body, &splitVideoInput); err != nil {
		return err
	}

	if err := media.SplitVideo(&splitVideoInput); err != nil {
		return err
	}
	log.Println("Successful media split: %s\n", string(body))
	return nil
}
