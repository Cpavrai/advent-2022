package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initCrates(initialisation []string) [][]string {
	crates := make([][]string, (len(initialisation[0])+1)/4)
	j := 0

	for r, row := range initialisation {
		j = 0
		if r != len(initialisation)-1 {
			for i := 1; i < len(row)-1; i += 4 {
				if row[i] != ' ' {
					crates[j] = append(crates[j], string(row[i]))
				}
				j += 1
			}
		}
	}
	return crates
}

func movements(match []string, crates [][]string) {
	size, _ := strconv.Atoi(match[0])
	source, _ := strconv.Atoi(match[1])
	destination, _ := strconv.Atoi(match[2])
	for i := 0; i < size; i++ {
		crates[destination-1] = append(crates[destination-1], "0")
		copy(crates[destination-1][1:], crates[destination-1])
		crates[destination-1][0] = crates[source-1][0]
		crates[source-1] = crates[source-1][1:]
	}
}

func computeHeads(initialisation, moves []string) string {
	crates := initCrates(initialisation)
	var re = regexp.MustCompile(`([0-9]+)`)
	result := ""

	for _, move := range moves {
		movements(re.FindAllString(move, -1), crates)
	}
	for _, crate := range crates {
		result += string(crate[0][0])
	}
	return result
}

func main() {
	if len(os.Args) == 2 {
		dat, err := os.ReadFile(os.Args[1])
		check(err)
		instructions := strings.Split(string(dat), "\n\n")
		if len(instructions) == 0 {
			fmt.Println("File empty.")
		} else if len(instructions) != 2 {
			fmt.Println("File unvalid.")
		} else {
			fmt.Println(computeHeads(strings.Split(instructions[0], "\n"), strings.Split(instructions[1], "\n")))
		}
	} else {
		fmt.Println("Error, you need to pass a file as argument.")
	}
}
