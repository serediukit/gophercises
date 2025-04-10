package utils

import "os"

func ReadFileToByte(filename string) ([]byte, error) {
	file, err := os.ReadFile("res/" + filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
