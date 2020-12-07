package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	tix := parseInput()
	seats := tixToSeats(tix)
	fmt.Println("Part 1: ", maxSeatId(seats))
	fmt.Println("Part 2: ", part2(seats))

}

func parseInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())

	}
	return result
}

func tixToSeats(tickets []string) [][2]int {
	var seats [][2]int

	for _, tix := range tickets {
		seats = append(seats, [2]int{backFront(tix[:7]), leftRight(tix[7:])})
	}
	return seats
}

func backFront(s string) int {
	max := 127.0
	min := 0.0
	for _, r := range s {
		if r == 'F' {
			max = math.Floor((max + min) / 2.0)
		} else if r == 'B' {
			min = math.Ceil((max + min) / 2.0)
		} else {
			log.Fatalf("Could not parse rune: %v\n", r)
		}
	}
	return int(max)
}

func leftRight(s string) int {
	max := 7.0
	min := 0.0
	for _, r := range s {
		if r == 'R' {
			min = math.Ceil((max + min) / 2)
		} else if r == 'L' {
			max = math.Floor((max + min) / 2)
		} else {
			log.Fatalf("Could not parse rune: %v\n", r)
		}
	}
	return int(min)
}

func seatId(s [2]int) int {
	return (s[0] * 8) + s[1]
}

func maxSeatId(seats [][2]int) int {
	max := -1
	for _, s := range seats {
		id := seatId(s)
		if id > max {
			max = id
		}
	}
	return max
}

func part2(seats [][2]int) int {
	var seatIds []int
	for _, s := range seats {
		seatIds = append(seatIds, seatId(s))
	}
	sort.Ints(seatIds)
	for i, s := range seatIds {
		if i == 0 || i == len(seatIds) {
			continue
		}
		if seatIds[i-1] != s-1 {
			return s - 1
		} else if seatIds[i+1] != s+1 {
			return s + 1
		}
	}
	return -1
}
