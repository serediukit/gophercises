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
	filePath = flag.String("filePath", "problems.csv", "The path to file where questions are containing")
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

func main() {
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

	questions := make([]question, len(records))
	for i, record := range records {
		questions[i] = question{
			id:       i + 1,
			question: record[0],
			answer:   strings.ToLower(strings.TrimSpace(record[1])),
		}
	}

	if *shuffle {
		shuffleQuestions(questions)
	}

	scanner := bufio.NewScanner(os.Stdin)
	countCorrect, countTotal := 0, len(records)

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

	fmt.Printf("Correct Answers is %d out of %d\n", countCorrect, countTotal)
}
