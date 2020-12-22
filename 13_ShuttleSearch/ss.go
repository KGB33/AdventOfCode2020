package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func nextBus(currTime int, busses []int) (int, int) {
	/*
		Given a current time stamp and a list of busses,
		returns the bus ID and the wait time.
	*/
	min := 9*10 ^ 45
	var bussId int
	for _, b := range busses {
		waitTime := b - currTime%b
		if waitTime < min && waitTime >= 0 {
			min = waitTime
			bussId = b
		}
	}
	return bussId, min
}

func prizeTime(B []int) int {
	/*
		Finds the first time that shuttle one and shuttle two
		are at the correct times, then it jumps to the next time
		that the two are in sync. Once it finds when shuttle 1, 2, and 3
		are in sync it jumps to the next time that occurs. Until all
		shuttles are in sync.
	*/
	b, diff := prizeData(B)
	jump := b[0]
	n := 1
	for i := 0; true; i += jump {
		if (i+diff[n])%b[n] == 0 {
			jump = jump * b[n]
			n++
		}
		if n >= len(b) {
			return i
		}
	}
	return -1
}

func prizeData(B []int) ([]int, []int) {
	var bOut, iOut []int
	for i, b := range B {
		if b == -1 {
			continue
		}
		bOut = append(bOut, b)
		iOut = append(iOut, i)
	}
	return bOut, iOut
}

func all(in []bool) bool {
	for _, b := range in {
		if !b {
			return false
		}
	}
	return true
}

func main() {
	currTime, bIds := parseInput()
	nextBusId, waitTime := nextBus(currTime, bIds)
	fmt.Printf("You have to wait %d mins for bus %d.\n\tThe solution for part 1 is: %d\n",
		waitTime, nextBusId, waitTime*nextBusId)
	fmt.Println("The answer to the Shuttle Compaine's riddle is:", prizeTime(bIds))
}

func parseInput() (int, []int) {
	file, err := os.Open("input.prod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	currTime, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	var bIds []int
	for _, b := range strings.Split(scanner.Text(), ",") {
		bId, err := strconv.Atoi(b)
		if err != nil {
			bId = -1
		}
		bIds = append(bIds, bId)
	}
	return currTime, bIds
}
