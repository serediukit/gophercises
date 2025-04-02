package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	filePath = "problems.csv"
)

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Cannot close file:", err)
		}
	}(file)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Cannot read csv file:", err)
	}

	fmt.Println(records)
}
