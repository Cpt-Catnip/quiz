package main

import (
	"encoding/csv"
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

	// read quiz
	quizProblems, err := loadQuiz(DefaultQuiz)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", quizProblems)

	// for every question/answer pair
	// ask for user input
	// compare input to answer

	// display results
}
