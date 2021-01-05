package main

import "fmt"

func getToN(in []int, bound int) int {
	turnNum := len(in)
	history := make(map[int][2]int)
	for i, v := range in {
		history[v] = [2]int{i + 1, 0}
	}
	lastSpoken := in[len(in)-1]
	for turnNum != bound {
		turnNum++
		// If its never been spoken before
		lsh := history[lastSpoken]
		if lsh[1] == 0 {
			lastSpoken = 0
		} else { // If its been said before
			lastSpoken = lsh[0] - lsh[1]
		}
		// Rewrite History
		history[lastSpoken] = [2]int{turnNum, history[lastSpoken][0]}
	}
	return lastSpoken
}

func main() {
	// INPUT := []int{0,3,6} // test
	INPUT := []int{13, 0, 10, 12, 1, 5, 8} // prod

	p1 := getToN(INPUT, 2020)
	p2 := getToN(INPUT, 30000000)
	fmt.Println("Part 1", p1)
	fmt.Println("Part 2", p2)

}
