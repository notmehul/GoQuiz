package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	csvFile, err := os.Open("problems.csv")

	if err != nil {
		log.Fatalln("Couldn't load the questions", err)
	}

	qna := csv.NewReader(csvFile)
	for {
		problem, err := qna.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var x string
		fmt.Printf("Question: %s ", problem[0])
		fmt.Scan(x)

		if x == problem[1] {
			fmt.Printf("Good Job")
		}
	}
}
