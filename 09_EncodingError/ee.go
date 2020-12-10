package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(code []int) int {
	preamble := code[:25]
	code = code[25:]
	var d int

	for len(code) > 0 {
		d, code = code[0], code[1:]
		if !isValid(d, preamble) {
			return d
		}
		_, preamble = preamble[0], append(preamble[1:], d)
	}
	return -1
}

func part2(code []int, goal int) []int {
	/*
		Create a range over the input splice,
		if the sum of the range is greater than
		the goal, remove the 1st element, if it
		is less than the goal, add the next
		element in the input
	*/
	var r []int
	r, code = code[:1], code[2:]
	for {
		s := sum(r)
		switch {
		case s < goal:
			r = append(r, code[0])
			_, code = code[0], code[1:]
		case s > goal:
			_, r = r[0], r[1:]
		case s == goal:
			return r
		}
	}
}

func sum(arr []int) int {
	var s int
	for _, d := range arr {
		s += d
	}
	return s
}

func max_min(arr []int) (int, int) {
	max, min := arr[0], arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		} else if i < min {
			min = i
		}
	}
	return max, min
}

func isValid(n int, p []int) bool {
	for _, a := range p {
		for _, b := range p {
			if a == b {
				continue
			} else if a+b == n {
				return true
			}
		}
	}
	return false
}

func main() {
	code := parseInput()
	invaid := part1(code)
	fmt.Println("Part 1: ", invaid)
	max, min := max_min(part2(code, invaid))
	fmt.Println("\n\nPart 2: ", max+min)
}

func parseInput() []int {
	file, err := os.Open("input")
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
	return out
}
