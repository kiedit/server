package main

import (
	"bytes"
	"log"
	"os/exec"

	user "kiedit/user"
	utils "kiedit/utils"
)

func main() {
	currentUser := new(user.User)
	currentUser.Init()

	flagsConfig := new(utils.FlagsConfig)
	flagsConfig.Init(currentUser.SessionDir)

	command := exec.Command("ffmpeg", "-i", flagsConfig.InputFile, "-f", "segment", "-segment_time", flagsConfig.Segment, "-c", "copy", flagsConfig.OutputDirPath)
	var stderr bytes.Buffer
	command.Stderr = &stderr
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
