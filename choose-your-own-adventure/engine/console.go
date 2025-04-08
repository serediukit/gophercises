package engine

import (
	"bufio"
	"cyoa/entity"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func Game(decoded map[string]*entity.Arc) {
	fmt.Println("Starting game ...")
	fmt.Println()

	currentArc := "intro"

GAME:
	for {
		arc := decoded[currentArc]

		printTitleAndStory(arc)

		if len(arc.Options) > 0 {
			printOptions(arc.Options)

			choice := getChoice(len(arc.Options) + 1)

			switch {
			case choice <= len(arc.Options):
				currentArc = arc.Options[choice-1].Arc
			default:
				break GAME
			}
		} else {
			fmt.Println("It's the END.")
			fmt.Println()
			break GAME
		}
	}
}

func printTitleAndStory(arc *entity.Arc) {
	fmt.Println()
	fmt.Println(arc.Title)
	fmt.Println()
	for _, story := range arc.Story {
		fmt.Println(story)
	}
	fmt.Println()
}

func printOptions(options []*entity.Option) {
	fmt.Println("What do you want to do?")
	fmt.Println()
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option.Text)
	}
	fmt.Printf("%d. Exit\n", len(options)+1)
	fmt.Println()
	fmt.Print("Your choice: ")
}

func getChoice(maxNum int) int {
	var choice int

	for {
		scanner.Scan()
		input := scanner.Text()
		c, err := strconv.Atoi(input)
		if err != nil || c > maxNum || c <= 0 {
			fmt.Printf("Invalid choice. Your input must be a number between 1 and %d.\n", maxNum)
		} else {
			choice = c
			break
		}
	}

	return choice
}
