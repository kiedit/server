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

	if err := queue.Consume(); err != nil {
		log.Fatal(err)
	}
}
