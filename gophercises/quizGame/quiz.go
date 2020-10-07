package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in " +
		"the format of 'question,answer'")
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
	//fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
			fmt.Println("Correct!")
		} else {
			fmt.Println("incorrect!!!")
		}
	}
	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
}

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