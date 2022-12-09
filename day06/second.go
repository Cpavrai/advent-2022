package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkValid(sequence string) bool {
	tmp_sequence := strings.Split(sequence, "")
	sort.Strings(tmp_sequence)
	sequence = strings.Join(tmp_sequence, "")
	for i := 1; i < len(sequence); i++ {
		if sequence[i] == sequence[i-1] {
			return false
		}
	}
	return true
}

func findMarker(packet string) int {
	index := 14

	for !checkValid(packet[index-14 : index]) {
		index += 1
	}
	return index
}

func main() {
	if len(os.Args) == 2 {
		dat, err := os.ReadFile(os.Args[1])
		check(err)
		packets := strings.Split(string(dat), "\n")
		if len(packets) == 0 {
			fmt.Println("File empty.")
		} else if len(packets) != 1 {
			fmt.Println("File unvalid.")
		} else {
			fmt.Println(findMarker(packets[0]))
		}
	} else {
		fmt.Println("Error, you need to pass a file as argument.")
	}
}
