package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b := parseInput()
	fmt.Println(b.run())
	fmt.Println(mutateBootCode(b))

}

type BootCode struct {
	instructions []Instruction
	accumulator  int
	visited      map[int]bool
}

func (b *BootCode) run() (int, error) {
	j := 0
	for j != len(b.instructions) {

		// Check if we've seen this index before
		if b.visited[j] {
			return b.accumulator, fmt.Errorf("Visited same instruction twice (index: %d)", j)
		} else {
			b.visited[j] = true
		}

		// Logic
		switch i := b.instructions[j]; i.op {
		case "acc":
			b.accumulator += i.arg
			j++
		case "jmp":
			j += i.arg
		case "nop":
			j++
		}
	}
	return b.accumulator, nil
}

func NewBootCode(i []Instruction) BootCode {
	return BootCode{i, 0, make(map[int]bool)}
}

type Instruction struct {
	op  string
	arg int
}

func mutateBootCode(b BootCode) (int, error) {
	for j, i := range b.instructions {
		if i.op == "jmp" {
			new_i := append([]Instruction{}, b.instructions...)
			new_i[j] = Instruction{"nop", i.arg}
			new_b := NewBootCode(new_i)
			r, err := new_b.run()
			if err == nil {
				return r, nil
			}
		} else if i.op == "nop" {
			new_i := append([]Instruction{}, b.instructions...)
			new_i[j] = Instruction{"jmp", i.arg}
			new_b := NewBootCode(new_i)
			r, err := new_b.run()
			if err == nil {
				return r, nil
			}
		}
	}
	return -1, fmt.Errorf("Could not Mutate into valid code")
}

func parseInput() BootCode {
	file, err := os.Open("input.prod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var out []Instruction
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)
		arg, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, Instruction{fields[0], arg})
	}
	return NewBootCode(out)
}
