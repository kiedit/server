package main

import (
	"log"

	media "kiedit/media"
	user "kiedit/user"
	utils "kiedit/utils"
)

func main() {
	currentUser := new(user.User)
	currentUser.Init()

	flagsConfig := new(utils.FlagsConfig)
	flagsConfig.Init()

	var splitVideoInput = media.SplitVideoInput{
		InputFile:     flagsConfig.InputFile,
		Segment:       flagsConfig.Segment,
		OutputDirPath: currentUser.SessionDir + "/output%03d.mp4",
	}

	if err := media.SplitVideo(&splitVideoInput); err != nil {
		log.Fatal(err)
	}
}
