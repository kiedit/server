package utils

import "flag"

type FlagsConfig struct {
	InputFile     string
	Segment       string
	OutputDirPath string
}

func (self *FlagsConfig) Init(workDir string) {
	inputFilePathPtr := flag.String("i", "filePath", "a string")
	segmentPtr := flag.String("s", "segment", "a string")

	flag.Parse()

	self.InputFile = *inputFilePathPtr
	self.Segment = *segmentPtr
	self.OutputDirPath = workDir + "/output%03d.mp4"
}
