package main

import (
	"kiedit/queue"
	"log"
)

func main() {
	queue := new(queue.QueueStruct)
	if err := queue.Connect(); err != nil {
		log.Fatal(err)
	}
	defer queue.Close()
	log.Println("Successfully connected to RabbitMQ instance")

	if err := queue.Consume(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully consumed RabbitMQ instance")
}
