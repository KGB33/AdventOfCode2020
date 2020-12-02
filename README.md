# Advent of Code 2020

My solutions to [AdventOfCode2020](https://adventofcode.com/2020https://adventofcode.com/2020)

All solutions will be written in Go.


# Notes

## Files

A major part of this challenge is reading input files.
An excellent stack overflow response can be found [here](https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559)

```Go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("/path/to/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
```
