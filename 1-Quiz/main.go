package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var QuizList []Quiz
var UserAnsList []string

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type Quiz struct {
	Q string
	A string
}

func checkAns(qNo int, ans string) bool {
	if ans == QuizList[qNo].A {
		return true
	} else {
		return false
	}
}
func getScore(ansList *[]string) int {
	var score int
	for i, ans := range *ansList {
		if checkAns(i, ans) {
			score++
		}
	}
	return score
}
func main() {
	//go run main.go -csv="problems.csv"
	fmt.Println("Wellocome to Go Quiz App")

	quitionsFilePath := flag.String("csv", "problems.csv", "pass the csv file path throug the flag")
	flag.Parse()

	file, err := os.Open(*quitionsFilePath)
	if err != nil {
		exit(fmt.Sprintf("Getting the error to open the fie at paht %s", *quitionsFilePath))
	}

	// linesbytes,err := io.ReadAll(file) // it will rturn the
	/*
		[53 43 53 44 49 48 10 49 43 49 44 50 10 56 43 51 44 49 49 10 49 43 50 44 51 10 56 43 54 44 49 52 10 51 43 49 44 52 10 49 43 52 44 53 10 53 43 49 44 54 10 50 43 51 44 53 10 51 43 51 44 54 10 50 43 52 44 54 10 53 43 50 44 55 10]
	*/
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		exit("Getting error to read the fiel")
	}

	// fmt.Println(linesbytes) //[[5+5 10] [1+1 2] [8+3 11] [1+2 3] [8+6 14] [3+1 4] [1+4 5] [5+1 6] [2+3 5] [3+3 6] [2+4 6] [5+2 7]]
	for _, row := range lines {
		// fmt.Println(row[0])
		QuizList = append(QuizList, Quiz{
			Q: row[0],
			A: row[1],
		})
	}

	sc := bufio.NewReader(os.Stdin)

	for i, quize := range QuizList {
		fmt.Printf("Q%d) %v : \n", i, quize.Q)
		fmt.Print("Ans --> ")
		str, err := sc.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		ans := strings.Trim(str, "\n")
		UserAnsList = append(UserAnsList, ans)
	}

	fmt.Printf("Congrats your Score is %d out of %d \n", getScore(&UserAnsList), len(QuizList))

}
