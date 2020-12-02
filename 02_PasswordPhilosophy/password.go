package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	passwords := parseFile("input.txt")

	part1 := 0
	part2 := 0
	for _, pw := range passwords {
		if isValidPart1(pw) {
			part1++
		}
		if isValidPart2(pw) {
			part2++
		}

	}
	fmt.Printf("There are %d vaild passwords for Part 1\n", part1)
	fmt.Printf("There are %d vaild passwords for Part 2\n", part2)

}

type password struct {
	min  int
	max  int
	char string
	pw   string
}

func parseFile(name string) [1000]password {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [1000]password
	var i int

	splitOn := regexp.MustCompile(`-| |: `)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := splitOn.Split(line, -1)

		min, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		result[i] = password{min, max, split[2], split[3]}
		i++
	}
	return result
}

func isValidPart1(p password) bool {
	count := strings.Count(p.pw, p.char)
	return p.min <= count && count <= p.max
}

func isValidPart2(p password) bool {
	return (string(p.pw[p.min-1]) == p.char) != (string(p.pw[p.max-1]) == p.char)
}
