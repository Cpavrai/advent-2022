package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkInside(assignment string) bool {
	orders := strings.Split(assignment, ",")
	firstElf := strings.Split(orders[0], "-")
	secondElf := strings.Split(orders[1], "-")
	firstElf1, _ := strconv.Atoi(firstElf[0])
	firstElf2, _ := strconv.Atoi(firstElf[1])
	secondElf1, _ := strconv.Atoi(secondElf[0])
	secondElf2, _ := strconv.Atoi(secondElf[1])

	return ((firstElf1 <= secondElf1) && (firstElf2 >= secondElf2)) ||
		((firstElf1 >= secondElf1) && (firstElf2 <= secondElf2))
}

func computeScore(assignments []string) int {
	result := 0

	for _, assignment := range assignments {
		if checkInside(assignment) {
			result += 1
		}
	}
	return result
}

func main() {
	if len(os.Args) == 2 {
		dat, err := os.ReadFile(os.Args[1])
		check(err)
		assignments := strings.Split(string(dat), "\n")
		if len(assignments) == 0 {
			fmt.Println("File empty or unvalid")
		} else {
			fmt.Println(computeScore(assignments))
		}
	} else {
		fmt.Println("Error, you need to pass a file as argument.")
	}
}
