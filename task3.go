package main

import (
    "encoding/csv"
    "fmt"
    "os"
	"io"
	"strconv"
	"log"
	"strings"
)

func main() {

	allAnswers := 0
	correctAnswers := 0

    f, err := os.Open("quiz.csv")
	if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    csvReader := csv.NewReader(f)
	for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }

		numbers := strings.Fields(rec[0])
		result, err := strconv.Atoi(strings.Fields(rec[1])[0])
		
		numbersSeparated := strings.Split(numbers[0], "+")
		x, err := strconv.Atoi(numbersSeparated[0])
		y, err := strconv.Atoi(numbersSeparated[1])
	
		if(x + y == result){
			correctAnswers = correctAnswers + 1
		}	
		allAnswers = allAnswers + 1
    }
   
	fmt.Println("All questions: ", allAnswers)
	fmt.Println("Correct Answers: ", correctAnswers)
}