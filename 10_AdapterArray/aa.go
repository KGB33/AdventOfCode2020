package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func part1(in []int) map[int]int {
	out := map[int]int{1: 0, 2: 0, 3: 0}
	var pre, cur int
	for len(in) > 0 {
		switch pre, cur, in = cur, in[0], in[1:]; cur - pre {
		case 1:
			out[1]++
		case 2:
			out[2]++
		case 3:
			out[3]++
		}
	}
	return out

}

func part2(in []int) int {
	// Add the wall and the Device
	in = append([]int{0}, in...)
	in = append(in, in[len(in)-1]+3)

	pathsToIndex := make([]int, len(in))
	pathsToIndex[0] = 1

	for i, a := range in {
		if i == 0 {
			continue
		}
		pathsToIndex[i] = pathsToIndex[i-1]
		if i > 1 && a-in[i-2] <= 3 {
			pathsToIndex[i] += pathsToIndex[i-2]
		}
		if i > 2 && a-in[i-3] <= 3 {
			pathsToIndex[i] += pathsToIndex[i-3]
		}

	}
	return pathsToIndex[len(pathsToIndex)-1]

}

func main() {
	in := parseInput()
	diffs := part1(in)
	fmt.Println("Part 1: ", diffs[1]*diffs[3])
	p2 := part2(in)
	fmt.Println("Part 2: ", p2)
}

func parseInput() []int {
	file, err := os.Open("input.prod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		digit, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, digit)
	}
	sort.Ints(out)
	return append(out, out[len(out)-1]+3)

}
