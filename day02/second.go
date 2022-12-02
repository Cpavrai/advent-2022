package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var score = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var turnScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var translate = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var logical = "XYZ"

func computeTurn(first, second string) int {
	// If LOSE
	if second == "X" {
		logicalIndex := (strings.Index(logical, translate[first]) + 2) % 3
		return turnScore[logical[logicalIndex:logicalIndex+1]]
	}
	// If DRAW
	if second == "Y" {
		return turnScore[translate[first]]
	}
	// If WIN
	if second == "Z" {
		logicalIndex := (strings.Index(logical, translate[first]) + 1) % 3
		return turnScore[logical[logicalIndex:logicalIndex+1]]
	}
	return 0
}

func computeScore(turns []string) int {
	result := 0

	for _, value := range turns {
		turn := strings.Split(value, " ")
		result += computeTurn(turn[0], turn[1]) + score[turn[1]]
	}
	return result
}

func main() {
	if len(os.Args) == 2 {
		dat, err := os.ReadFile(os.Args[1])
		check(err)
		chars := strings.Split(string(dat), "\n")
		if len(chars) == 0 {
			fmt.Println("File empty or unvalid")
		} else {
			fmt.Println(computeScore(chars))
		}
	} else {
		fmt.Println("Error, you need to pass a file as argument.")
	}
}
