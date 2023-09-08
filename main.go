package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

const DefaultQuiz = "problems.csv"

type Questions [][]string

func loadQuiz(filename string) (Questions, error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return Questions{}, err
	}

	// load csv from file
	qs, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return Questions{}, err
	}
	return qs, nil
}

func main() {
	// read command line args
	filePathPtr := flag.String("file", DefaultQuiz, "filepath for quiz")

	// read quiz
	quizProblems, err := loadQuiz(*filePathPtr)
	if err != nil {
		log.Fatal(err)
	}

	nCorrect := 0

	for i, problem := range quizProblems {
		// prompt question
		fmt.Printf("Question %d: %s\n", i+1, problem[0])

		// get answer
		fmt.Print("your answer: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()

		// check answer
		if input == problem[1] {
			nCorrect++
		}
	}

	// display results
	fmt.Printf("Congrats on finishing! You answered %d/%d questions correct.\n", nCorrect, len(quizProblems))
}
