package utils

import "flag"

type FlagsConfig struct {
	InputFile string
	Segment   string
}

func (self *FlagsConfig) Init() {
	inputFilePathPtr := flag.String("i", "filePath", "a string")
	segmentPtr := flag.String("s", "segment", "a string")

	flag.Parse()

	self.InputFile = *inputFilePathPtr
	self.Segment = *segmentPtr
}
