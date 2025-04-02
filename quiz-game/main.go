package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	filePath = flag.String("filePath", "problems.csv", "The path to file where questions are containing")
	limit    = flag.Int("limit", 30, "The time in second that you will have to complete the quiz")
)

func main() {
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Cannot read file:", err)
		return
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
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	countCorrect, countTotal := 0, len(records)

	for i, record := range records {
		fmt.Printf("Problem #%d: %s = ", i+1, record[0])

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		answer := strings.TrimSpace(record[1])
		answer = strings.ToLower(answer)

		if answer == input {
			countCorrect++
		}
	}

	fmt.Printf("Correct Answers is %d out of %d\n", countCorrect, countTotal)
}
