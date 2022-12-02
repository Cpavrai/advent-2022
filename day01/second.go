package main

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func compute(elve string)(int) {
    result := 0

    for _, value := range strings.Split(elve, "\n") {
        intVar, err := strconv.Atoi(value)
        if (err != nil) {
            fmt.Println(err)
            return 0
        }
        result += intVar
    }
    return result
}

func getThreeBiggest(elves []string)(int) {
    sort.SliceStable(elves, func(i, j int) bool {
        return compute(elves[i]) > compute(elves[j])
    })

    return compute(elves[0]) + compute(elves[1]) + compute(elves[2])
}

func main() {
    if (len(os.Args) == 2) {
        dat, err := os.ReadFile(os.Args[1])
        check(err)
        chars := strings.Split(string(dat), "\n\n")
        if (len(chars) == 0) {
            fmt.Println("File empty or unvalid")
        } else {
            fmt.Println(getThreeBiggest(chars))
        }
    } else {
        fmt.Println("Error, you need to pass a file as argument.")
    }
}