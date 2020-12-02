package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	report := parseFile("input.txt")
	part1(report)
	part2(report)

}

func parseFile(filepath string) [200]int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var report [200]int // the file has 200 lines
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	return report
}

func part1(list [200]int) {
	for _, x := range list {
		for _, y := range list {
			if x+y == 2020 {
				fmt.Println("Part 1 Solution")
				fmt.Printf("\t%d * %d = %d\n", x, y, x*y)
				return
			}
		}

	}
}

func part2(list [200]int) {
	for _, x := range list {
		for _, y := range list {
			for _, z := range list {
				if x+y+z == 2020 {
					fmt.Println("Part 2 Solution")
					fmt.Printf("\t%d * %d * %d = %d\n", x, y, z, x*y*z)
					return
				}
			}
		}

	}
}
