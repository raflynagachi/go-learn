package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	var correct int
	var answer string

	csvFilename := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		fmt.Scanf("%s\n", &answer)
		checkProblem(p, answer, &correct)
	}

	fmt.Printf("Correct %d/%d\n", correct, len(lines))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func checkProblem(p problem, answer string, correct *int) {
	if answer == p.a {
		*correct++
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
