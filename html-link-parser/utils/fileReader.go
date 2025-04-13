package utils

import (
	"io"
	"os"
)

func ReaderFromFile(filename string) (io.Reader, error) {
	file, err := os.Open("res/" + filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
