package queue

import "log"

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

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Println("Received Message: %s\n", msg.Body)
		}
	}()

	log.Println("Waiting for messages...")
	<-forever

	return nil
}
