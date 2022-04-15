package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	var correct int
	var answer string

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("timeLimit", 30, "the time limit for the quiz in seconds")
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

	var menu string
	for strings.ToLower(menu) != "y" {
		fmt.Println("Ready for quiz? (y/n)")
		fmt.Scan(&menu)
	}
	fmt.Println("START")

	quizTimer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-quizTimer.C:
			fmt.Printf("\nCorrect %d/%d\n", correct, len(lines))
			return
		case <-answerCh:
			checkProblem(p, answer, &correct)
		}
	}
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
