package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Ferry struct {
	seats [][]rune // One of '#', 'L' or ''
}

func (f Ferry) CountSeats() map[rune]int {
	count := make(map[rune]int)
	for _, row := range f.seats {
		for _, v := range row {
			count[v]++
		}
	}
	return count
}

func (f Ferry) RunUntilStatic(depth int, tolerance int) Ferry {
	seatsChanged := -1
	for seatsChanged != 0 {
		fmt.Println(f)
		fmt.Println()
		f, seatsChanged = f.Run(depth, tolerance)
	}
	return f
}

func (f Ferry) Run(depth int, tolerance int) (Ferry, int) {
	/*
		Processes one Gereration of the game.
		Returns the number of seats changed
	*/
	seatsChanged := 0
	nextGeneration := make([][]rune, len(f.seats))
	for i := range nextGeneration {
		nextGeneration[i] = make([]rune, len(f.seats[0]))
	}
	for x, seatRow := range f.seats {
		for y, seat := range seatRow {
			// Ignore Floors
			if seat == 0 {
				nextGeneration[x][y] = 0
				continue
			}
			alive := f.findOccupiedSeats([]int{x, y}, depth)
			if seat == 'L' && alive == 0 {
				nextGeneration[x][y] = '#'
				seatsChanged++
			} else if seat == '#' && alive >= tolerance {
				nextGeneration[x][y] = 'L'
				seatsChanged++
			} else {
				nextGeneration[x][y] = seat
			}
		}
	}
	return Ferry{nextGeneration}, seatsChanged

}

func (f Ferry) findOccupiedSeats(seatIndex []int, distance int) int {
	/*
		This function checks for occupied seats in
		every direction up to the distance provided.
		If the distance == -1, then it checks until
		there is no next seat to check.
	*/
	if distance == -1 {
		distance = len(f.seats) * len(f.seats[0]) // Will be longer than the longest diagonal
	}
	numOccupiedSeats := 0
	scalar := 1
	vectors := [][]int{{1, 1}, {1, 0}, {-1, 1}, {-1, 0}, {1, -1}, {-1, -1}, {0, 1}, {0, -1}}
	for scalar <= distance && len(vectors) > 0 {
		var newVectors [][]int
		for _, v := range vectors {
			x, y := seatIndex[0]+v[0]*scalar, seatIndex[1]+v[1]*scalar
			if 0 > x || x >= len(f.seats) {
				continue
			}
			if 0 > y || y >= len(f.seats[0]) {
				continue
			}
			ajSeat := f.seats[x][y]

			if ajSeat == '#' {
				numOccupiedSeats++
			} else if ajSeat == 'L' {
				continue
			} else {
				newVectors = append(newVectors, v)
			}
		}
		scalar++
		vectors = newVectors
	}
	return numOccupiedSeats
}

func (f Ferry) String() string {
	out := ""
	for _, line := range f.seats {
		for _, seat := range line {
			out += string(seat)
		}
		out += "\n"
	}
	return out
}

func main() {
	f := parseInput()
	f = f.RunUntilStatic(-1, 5)
	fmt.Println("Part 1: ", f.CountSeats()['#'])
}

func parseInput() Ferry {
	file, err := os.Open("input.prod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var seats [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seats = append(seats, []rune(scanner.Text()))
	}
	return Ferry{seats}
}

func remove(s [][]int, i int) [][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
