package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Ferry struct {
	seats      [][]int
	seatStatus map[int]rune // One of '#', 'L' or ''
}

func (f Ferry) CountSeats() map[rune]int {
	count := make(map[rune]int)
	for _, v := range f.seatStatus {
		count[v]++
	}
	return count
}

func (f Ferry) RunUntilStatic() Ferry {
	seatsChanged := -1
	for seatsChanged != 0 {
		f, seatsChanged = f.Run()
	}
	return f
}

func (f Ferry) Run() (Ferry, int) {
	/*
		Processes one Gereration of the game.
		Returns the number of seats changed
	*/
	seatsChanged := 0
	nextGeneration := make(map[int]rune)
	for seat, neighbors := range f.seats {
		// Ignore Floors
		if f.seatStatus[seat] == 0 {
			nextGeneration[seat] = 0
			continue
		}
		var alive, dead int
		for _, neighbor := range neighbors {
			switch f.seatStatus[neighbor] {
			case 0:
				continue
			case 'L':
				dead++
			case '#':
				alive++
			}
		}
		if f.seatStatus[seat] == 'L' && alive == 0 {
			nextGeneration[seat] = '#'
			seatsChanged++
		} else if f.seatStatus[seat] == '#' && alive >= 4 {
			nextGeneration[seat] = 'L'
			seatsChanged++
		} else {
			nextGeneration[seat] = f.seatStatus[seat]
		}
	}
	return Ferry{f.seats, nextGeneration}, seatsChanged

}

func (f Ferry) String() string {
	out := ""
	var max int
	for _, v := range f.seats[0] {
		if v > max {
			max = v
		}
	}
	lineLength := max - 1

	for i := 0; i < len(f.seatStatus); i++ {
		v := f.seatStatus[i]
		if i%lineLength == 0 {
			out += "\n"
		}
		if v == 0 {
			out += "."
		} else {
			out += string(v)
		}

	}
	return out
}

func main() {
	f := parseInput()
	f = f.RunUntilStatic()
	fmt.Println("Part 1: ", f.CountSeats()['#'])
}

func parseInput() Ferry {
	file, err := os.Open("input.prod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var seats [][]int
	seatStatus := make(map[int]rune)
	scanner := bufio.NewScanner(file)
	var lineNumber int
	for scanner.Scan() {
		text := "." + scanner.Text() + "."
		lineLength := len(text)
		for i, r := range text {
			seatIndex := i + lineNumber*lineLength
			if r == '.' {
				seatStatus[seatIndex] = 0
			} else {
				seatStatus[seatIndex] = r
			}
			seats = append(seats, []int{
				// Previous Line
				i - 1 + (lineNumber-1)*lineLength,
				i + (lineNumber-1)*lineLength,
				i + 1 + (lineNumber-1)*lineLength,
				// Same Line
				seatIndex - 1,
				seatIndex + 1,
				// Next Line
				i - 1 + (lineNumber+1)*lineLength,
				i + (lineNumber+1)*lineLength,
				i + 1 + (lineNumber+1)*lineLength,
			})

		}
		lineNumber++
	}
	return Ferry{seats, seatStatus}
}
