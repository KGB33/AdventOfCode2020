// rewrote 3d part 1 code to work for 4d
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// const DIMENSION = 3 // Part 1
const DIMENSION = 4 // Part 2

func main() {
	pocketDimension := parseInput("input.prod")
	fmt.Println(pocketDimension)
	for range make([]int, 6) {
		pocketDimension.Cycle()
	}
	fmt.Printf("Part %d solution: %d\n", DIMENSION-2, len(pocketDimension.grid))
}

func parseInput(fileName string) Dimension {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open '%s' - Err: %s", fileName, err)
	}
	defer file.Close()

	output := make(map[[DIMENSION]int]Cube)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		text := scanner.Text()
		for x, r := range text {
			if r == '#' {
				output[[DIMENSION]int{x, y}] = Cube{true, 0}
			}
		}
		y++
	}
	return Dimension{output}
}

func generateNeighborCoords(coords [DIMENSION]int) [][DIMENSION]int {
	output := [][DIMENSION]int{}
	deltaPermutations := func() [][DIMENSION]int {
		permutations := [][DIMENSION]int{}
		currPermutation := [DIMENSION]int{}
		vals := []int{-1, 0, 1}
		pn := make([]int, DIMENSION)
		k := len(vals)
		for {
			for i, x := range pn {
				currPermutation[i] = vals[x]
			}
			permutations = append(permutations, currPermutation)

			// increment permutation number
			for i := 0; ; {
				pn[i]++
				if pn[i] < k {
					break
				}
				pn[i] = 0
				i++
				if i == DIMENSION {
					return permutations
				}
			}
		}
	}()
	for _, delta := range deltaPermutations {
		if delta == [DIMENSION]int{} {
			continue
		}
		newNeighbor := coords
		for i := 0; i < DIMENSION; i++ {
			newNeighbor[i] = coords[i] + delta[i]
		}
		output = append(output, newNeighbor)
	}
	return output
}

type Dimension struct {
	// Using a map allows us to have negative indexes, whereas a 3d splice does not
	grid map[[DIMENSION]int]Cube
}

func (d *Dimension) CalculateNeighbors() {
	d.ClearNeighbors()
	var keys [][DIMENSION]int
	for k := range d.grid {
		keys = append(keys, k)
	}

	for _, coords := range keys {
		for _, neighborCoords := range generateNeighborCoords(coords) {
			if old_cube, ok := d.grid[neighborCoords]; ok {
				d.grid[neighborCoords] = Cube{old_cube.status, old_cube.active_neighbors + 1}
			} else {
				d.grid[neighborCoords] = Cube{false, 1}
			}
		}
	}
}

func (d *Dimension) ClearNeighbors() {
	var keys [][DIMENSION]int
	for k := range d.grid {
		keys = append(keys, k)
	}

	for _, k := range keys {
		d.grid[k] = Cube{d.grid[k].status, 0}
	}
}

func (d *Dimension) Cycle() {
	d.CalculateNeighbors()
	d.CalculateNextState()
	d.PruneDeadCubes()
}

func (d *Dimension) CalculateNextState() {
	g_loop := d.grid
	for coords, cube := range g_loop {
		if cube.active_neighbors >= 2 {
		}
		// If a cube is active and exactly 2 or 3 of its neighbors are also active,
		// the cube remains active. Otherwise, the cube becomes inactive.
		if cube.status {
			if !(cube.active_neighbors == 2 || cube.active_neighbors == 3) {
				d.grid[coords] = Cube{false, 0}
			}
		} else {

			// If a cube is inactive but exactly 3 of its neighbors are active,
			// the cube becomes active. Otherwise, the cube remains inactive.
			if cube.active_neighbors == 3 {
				d.grid[coords] = Cube{true, 0}
			}
		}
	}
}

func (d *Dimension) PruneDeadCubes() {
	g_loop := d.grid
	for coords, cube := range g_loop {
		if !cube.status {
			delete(d.grid, coords)
		}
	}
}

func (d Dimension) String() string {
	output := ""
	for cords, cube := range d.grid {
		output += fmt.Sprintf("%v -- %v\n", cords, cube)
	}
	return output
}

type Cube struct {
	status           bool
	active_neighbors int
}

func (c Cube) String() string {
	if c.status {
		return "#"
	}
	return "."
}
