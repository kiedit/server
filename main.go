package main

import (
	"bytes"
	"flag"
	"fmt"
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

	err := os.Mkdir(sessionDir, 0755)

	return err
}

func main() {
	inputFilePathPtr := flag.String("i", "filePath", "a string")
	segmentPtr := flag.String("s", "segment", "a string")

	userSessionId := generateUserSessionID()

	if err := createSessionDirectory(userSessionId); err != nil {
		panic(err)
	}

	flag.Parse()

	inputFile := *inputFilePathPtr
	segment := *segmentPtr
	outputDirPath := "./dist/" + userSessionId + "/output%03d.mp4"

	command := exec.Command("ffmpeg", "-i", inputFile, "-f", "segment", "-segment_time", segment, "-c", "copy", outputDirPath)
	var stderr bytes.Buffer
	command.Stderr = &stderr

	fmt.Println(command.String())

	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
