package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	uuid "github.com/satori/go.uuid"
)

func generateUserSessionID() string {
	return uuid.NewV4().String()
}

func createSessionDirectory(sessionID string) error {
	sessionDir := "./dist/" + sessionID

	return os.Mkdir(sessionDir, 0755)
}

func main() {
	userSessionId := generateUserSessionID()

	command := exec.Command("sh", "run.sh")
	var stderr bytes.Buffer
	command.Stderr = &stderr

	if err := createSessionDirectory(userSessionId); err != nil {
		log.Fatal(err)
	}

	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
