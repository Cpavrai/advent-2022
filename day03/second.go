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

var score = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func findDuplicate(first, second, third string) string {
	for _, letter := range first {
		if strings.Index(second, string(letter)) != -1 &&
			strings.Index(third, string(letter)) != -1 {
			return string(letter)
		}
	}
	return string(first[0])
}

func computeScore(rucksacks []string) int {
	result := 0

	for index := 0; index < len(rucksacks); index += 3 {
		result += 1 + strings.Index(
			score,
			findDuplicate(
				rucksacks[index],
				rucksacks[index+1],
				rucksacks[index+2],
			),
		)
	}
	return result
}

func main() {
	if len(os.Args) == 2 {
		dat, err := os.ReadFile(os.Args[1])
		check(err)
		rucksacks := strings.Split(string(dat), "\n")
		if len(rucksacks) == 0 {
			fmt.Println("File empty or unvalid")
		} else if len(rucksacks)%3 != 0 {
			fmt.Println("File not well formatted")
		} else {
			fmt.Println(computeScore(rucksacks))
		}
	} else {
		fmt.Println("Error, you need to pass a file as argument.")
	}
}
