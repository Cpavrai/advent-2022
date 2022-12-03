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

func findDuplicate(rucksack string) string {
	rucksackLength := len(rucksack)
	first := rucksack[:(rucksackLength / 2)]
	second := rucksack[(rucksackLength / 2):]
	for _, letter := range first {
		if strings.Index(second, string(letter)) != -1 {
			return string(letter)
		}
	}
	return string(first[0])
}

func computeScore(rucksacks []string) int {
	result := 0

	for _, rucksack := range rucksacks {
		result += strings.Index(score, findDuplicate(rucksack)) + 1
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
		} else {
			fmt.Println(computeScore(rucksacks))
		}
	} else {
		fmt.Println("Error, you need to pass a file as argument.")
	}
}
