package main

import (
    "fmt"
    "strings"
    "strconv"
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

func getBiggest(elves []string)(int) {
    result := compute(elves[0])

    for _, value := range elves {
        tmp := compute(value)
        if (tmp > result) {
            result = tmp
        }
    }
    return result
}

func main() {
    if (len(os.Args) == 2) {
        dat, err := os.ReadFile(os.Args[1])
        check(err)
        chars := strings.Split(string(dat), "\n\n")
        if (len(chars) == 0) {
            fmt.Println("File empty or unvalid")
        } else {
            fmt.Println(getBiggest(chars))
        }
    } else {
        fmt.Println("Error, you need to pass a file as argument.")
    }
}