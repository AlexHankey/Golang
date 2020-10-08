// Quiz game which tests your simple maths
// problems which you will be timed on as well

// Any changes you need to run
// go build .
// first

// Test the time limit = go run quiz.go -Limit=30
// Test the csv file = go run quiz.go -csv=problems.csv
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)


func parseLines(lines [][]string) []problem {
	prob := make([]problem, len(lines))
	for i, line := range lines {
		prob[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return prob
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in " +
		"the format of 'question,answer'")
	timeLimit := flag.Int("Limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <- timer.C:
			fmt.Println()
			break problemLoop
		case answer := <- answerChannel:
			if answer == p.a {
				correct++
				fmt.Println("Correct!")
			} else {
				fmt.Println("incorrect!")
			}
		}
	}
	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
}