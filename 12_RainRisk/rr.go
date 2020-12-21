package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func travelBoat(ins []rune, vals []int, a int, b int) (int, int) {
	var x, y int
	for j, i := range ins {
		switch i {
		case 'N':
			b += vals[j]
		case 'S':
			b -= vals[j]
		case 'E':
			a += vals[j]
		case 'W':
			a -= vals[j]
		case 'L':
			a, b = rotateVector(a, b, vals[j])
		case 'R':
			a, b = rotateVector(a, b, -vals[j])
		case 'F':
			x += a * vals[j]
			y += b * vals[j]
		}
	}

	return x, y
}

func rotateVector(a int, b int, theta int) (int, int) {
	u := a*degreeCos(theta) - b*degreeSin(theta)
	v := a*degreeSin(theta) + b*degreeCos(theta)
	return u, v
}

func degreeCos(d int) int {
	return int(math.Cos(float64(d) * math.Pi / 180))
}

func degreeSin(d int) int {
	return int(math.Sin(float64(d) * math.Pi / 180))
}

func main() {
	ins, vals := parseInput()
	x, y := travelBoat(ins, vals, 10, 1)
	fmt.Printf("The boat traveled to: (%d, %d), a distance of %d\n",
		x, y, int(math.Abs(float64(x)))+int(math.Abs(float64(y))))
}

func parseInput() ([]rune, []int) {
	file, err := os.Open("input.prod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []rune
	var values []int
	for scanner.Scan() {
		text := scanner.Text()
		instruction, value := rune(text[0]), text[1:]
		val, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, instruction)
		values = append(values, val)
	}
	return instructions, values

}
