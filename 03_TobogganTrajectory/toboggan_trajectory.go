package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	trees := parseInput()
	part1 := checkSlope(3, 1, trees)
	fmt.Printf("Part 1: %d\n\n", part1)

	r1d1 := checkSlope(1, 1, trees)
	r3d1 := part1
	r5d1 := checkSlope(5, 1, trees)
	r7d1 := checkSlope(7, 1, trees)
	r1d2 := checkSlope(1, 2, trees)

	fmt.Print("Part 2:\n")
	fmt.Printf("\tr1d1: %d\n\tr3d1: %d\n\tr5d1: %d\n\tr7d1: %d\n\tr1d2: %d", r1d1, r3d1, r5d1, r7d1, r1d2)
	fmt.Printf("\n\n\tSolution: %d", r1d1*r3d1*r5d1*r7d1*r1d2)
}

func checkSlope(run int, rise int, trees [][]bool) int {
	var x, y, total int
	for {
		// Check if Tree
		if trees[y][x] {
			total++
		}
		// Increment x & y
		y += rise
		x = (x + run) % 31

		if y >= 323 { // number of rows
			return total
		}

	}
}

func parseInput() [][]bool {
	// create array
	result := make([][]bool, 323)
	for i := range result {
		result[i] = make([]bool, 31)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var y int
	for scanner.Scan() {
		for x, r := range scanner.Text() {
			result[y][x] = (r == '#')
		}
		y++

	}

	return result

}
