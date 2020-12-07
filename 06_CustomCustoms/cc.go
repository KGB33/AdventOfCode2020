package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	groups := parseInput()
	fmt.Println("Part 1: ", part1(groups))
	fmt.Println("Part 2: ", part2(groups))
	fmt.Println(' ')
}

func part1(input []string) int {
	gs := inputToSet(input)
	var total int
	for _, g := range gs {
		total += len(g)
	}
	return total
}

func part2(input []string) int {
	var total int
	for _, group := range input {
		numPeople := len(strings.Fields(group))
		for _, v := range groupCounter(group) {
			if v == numPeople {
				total++
			}
		}
	}
	return total
}

func groupCounter(s string) map[rune]int {
	group := make(map[rune]int)
	for _, r := range s {
		group[r]++
	}
	delete(group, ' ')
	fmt.Printf("%s --> %v\n", s, group)
	return group
}

func inputToSet(in []string) []map[rune]bool {
	var output []map[rune]bool
	group := make(map[rune]bool)
	for _, g := range in {
		for _, r := range g {
			group[r] = true
		}
		delete(group, ' ')
		output = append(output, group)
		group = make(map[rune]bool)
	}
	return output

}

func parseInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string
	group := ""

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			output = append(output, group)
			group = ""
			continue
		}
		group += text + " "
	}
	return output
}
