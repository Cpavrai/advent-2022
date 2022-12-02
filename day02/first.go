package main

import (
    "fmt"
    "strings"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var score = map[string]int{
    "X": 1,
    "Y": 2,
    "Z": 3,
}

var translate = map[string]string{
    "X": "A",
    "Y": "B",
    "Z": "C",
}

func computeTurn(first, second string)(int) {
    // If DRAW
    if (first == translate[second]) {
        return 3
    }
    // If WIN
    if ((first == "A" && translate[second] == "B") ||
        (first == "B" && translate[second] == "C") ||
        (first == "C" && translate[second] == "A")) {
        return 6
    }
    return 0
}

func computeScore(turns []string)(int) {
    result := 0

    for _, value := range turns {
        turn := strings.Split(value, " ")
        result += computeTurn(turn[0], turn[1]) + score[turn[1]]
    }
    return result
}

func main() {
    if (len(os.Args) == 2) {
        dat, err := os.ReadFile(os.Args[1])
        check(err)
        chars := strings.Split(string(dat), "\n")
        if (len(chars) == 0) {
            fmt.Println("File empty or unvalid")
        } else {
            fmt.Println(computeScore(chars))
        }
    } else {
        fmt.Println("Error, you need to pass a file as argument.")
    }
}