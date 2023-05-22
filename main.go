package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	command := exec.Command("sh", "run.sh")

	var stderr bytes.Buffer
	command.Stderr = &stderr

	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
