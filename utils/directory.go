package utils

import (
	"os"
	"strings"
)

type Directory struct {
}

func (self *Directory) CreateDirectory(path []string) (string, error) {
	dir := strings.Join(path, "/")
	err := os.MkdirAll(dir, 0755)

	return dir, err
}
