package media

import (
	"bytes"
	"os/exec"
)

type SplitVideoInput struct {
	InputFile     string
	Segment       string
	OutputDirPath string
}

func SplitVideo(input *SplitVideoInput) error {
	command := exec.Command("ffmpeg", "-i", input.InputFile, "-f", "segment", "-segment_time", input.Segment, "-c", "copy", input.OutputDirPath)
	var stderr bytes.Buffer
	command.Stderr = &stderr

	return command.Run()
}
