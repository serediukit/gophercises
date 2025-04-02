package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	filePath = flag.String("filepath", "problems.csv", "The path to file where questions are containing")
	limit    = flag.Int("limit", 30, "The time in second that you will have to complete the quiz")
	shuffle  = flag.Bool("shuffle", false, "If true - the questions will appear in random order")
)

type question struct {
	id       int
	question string
	answer   string
}

func init() {
	flag.Parse()
}

func shuffleQuestions(a []question) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func newCsvReader(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Cannot read file:", err)
		os.Exit(1)
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
		os.Exit(1)
	}

	return records
}

func getQuestions(records [][]string) []question {

	questions := make([]question, len(records))
	for i, record := range records {
		questions[i] = question{
			id:       i + 1,
			question: record[0],
			answer:   strings.ToLower(strings.TrimSpace(record[1])),
		}
	}

	return questions
}

func quiz(questions []question) int {
	scanner := bufio.NewScanner(os.Stdin)
	countCorrect := 0

	fmt.Println("As you will be ready - press Enter to start...")
	scanner.Scan()

	timer := time.After(time.Duration(*limit) * time.Second)

	doneCh := make(chan struct{})

	go func() {
		for i, q := range questions {
			fmt.Printf("Problem #%d: %s = ", i+1, q.question)

			scanner.Scan()
			input := scanner.Text()
			input = strings.TrimSpace(input)
			input = strings.ToLower(input)

			if q.answer == input {
				countCorrect++
			}
		}
		doneCh <- struct{}{}
	}()

	select {
	case <-timer:
		fmt.Println("\nTime out")
	case <-doneCh:
	}

	return countCorrect
}

func main() {
	records := newCsvReader(*filePath)
	questions := getQuestions(records)

	if *shuffle {
		shuffleQuestions(questions)
	}

	countCorrect := quiz(questions)

	fmt.Printf("Correct Answers is %d out of %d (%.0f%%)\n", countCorrect, len(questions), float32(countCorrect)/float32(len(questions))*100)
}
