package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Ticket Translation")
	rules, your_ticket, nearby_tickets := read_data("input.prod")
	fmt.Printf("Solution to Part 1: %d\n", part1(rules, nearby_tickets))
	fmt.Printf("Solution to Part 2: %d\n", part2(rules, nearby_tickets, your_ticket))
}

// Find the sum of the values that are invalid for every rule
func part1(rules []Rule, nearby_tickets [][]int) int {
	applied := [][][]bool{}
	for _, r := range rules {
		applied = append(applied, r.Apply(nearby_tickets))
	}

	reduced := Reduce(applied)

	sum := 0
	for i, row := range reduced {
		for j, val := range row {
			if !val {
				sum += nearby_tickets[i][j]
			}
		}
	}
	return sum
}

func part2(rules []Rule, nearby_tickets [][]int, your_ticket []int) int {
	nearby_tickets = pruneInvaidTickets(nearby_tickets, rules)

	applied := [][][]bool{} // Rule - Ticket - Field
	for _, r := range rules {
		applied = append(applied, r.Apply(nearby_tickets))
	}

	possRules := map[int][]string{}
	for i, ruleMatrix := range applied {
		for j, field := range transposeBool(ruleMatrix) {
			if all(field) {
				// fmt.Printf("%dth Col could be %s\n", j, rules[i].name)
				possRules[j] = append(possRules[j], rules[i].name)
			}
		}
	}
	output := 1
	for k, v := range ReduceMap(possRules) {
		if strings.Contains(v, "departure") {
			output *= your_ticket[k]
		}
	}
	return output
}

func pruneInvaidTickets(tix [][]int, rules []Rule) [][]int {
	applied := [][][]bool{}
	for _, r := range rules {
		applied = append(applied, r.Apply(tix))
	}
	validMatrix := Reduce(applied)
	validTix := [][]int{}
	for i, row := range validMatrix {
		sentinel := true
		for _, val := range row {
			if !val {
				sentinel = val
				break
			}
		}
		if sentinel {
			validTix = append(validTix, tix[i])
		}
	}
	return validTix
}

func read_data(f string) ([]Rule, []int, [][]int) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("Could not open %s, err: %s\n", f, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	section := 0
	rules := []Rule{}
	var your_ticket []int
	nearby_tickets := [][]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			section++
			scanner.Scan()
			continue
		}
		switch section {
		case 0: // Rules
			rules = append(rules, parse_rule(text))
		case 1: // Your Ticket
			your_ticket = parse_ticket(text)
		case 2: // Nearby Tickets
			nearby_tickets = append(nearby_tickets, parse_ticket(text))
		default:
			log.Fatalf("Unexpected section: %d (text: %s)\n", section, text)
		}

	}
	return rules, your_ticket, nearby_tickets
}

func parse_rule(input string) Rule {
	splitOn := regexp.MustCompile(`: |-| or `)
	vals := splitOn.Split(input, -1) // [name, L1, L2, U1, U2]

	L1, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatalf("Could not parse %s to int (split from %s) - Err: %s\n", vals[1], input, err)
	}

	L2, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Fatalf("Could not parse %s to int (split from %s) - Err: %s\n", vals[2], input, err)
	}

	U1, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Fatalf("Could not parse %s to int (split from %s) - Err: %s\n", vals[3], input, err)
	}

	U2, err := strconv.Atoi(vals[4])
	if err != nil {
		log.Fatalf("Could not parse %s to int (split from %s) - Err: %s\n", vals[4], input, err)
	}

	return Rule{vals[0], [2]int{L1, L2}, [2]int{U1, U2}}
}

func parse_ticket(input string) []int {
	output := []int{}
	for _, s := range strings.Split(input, ",") {
		d, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not parse %s to int - Err: %s\n", s, err)
		}
		output = append(output, d)
	}
	return output
}

// Tickets are []ints

type Rule struct {
	name        string
	lower_range [2]int
	upper_range [2]int
}

func (r Rule) validate(num int) bool {
	return (r.lower_range[0] <= num && num <= r.lower_range[1]) ||
		(r.upper_range[0] <= num && num <= r.upper_range[1])
}

// Apply the rule to every value in the provided matrex
func (r Rule) Apply(tixMatrix [][]int) [][]bool {
	output := make([][]bool, len(tixMatrix))
	for r := range output {
		output[r] = make([]bool, len(tixMatrix[r]))
	}

	for i, row := range tixMatrix {
		for j, val := range row {
			output[i][j] = r.validate(val)
		}
	}

	return output
}

func Reduce(input [][][]bool) [][]bool {
	output := make([][]bool, len(input[0]))
	for r := range output {
		output[r] = make([]bool, len(input[0][r]))
	}

	for _, TxF := range input {
		for i, tix := range TxF {
			for j, val := range tix {
				if val {
					output[i][j] = val
				}
			}
		}
	}
	return output
}

func all(input []bool) bool {
	for _, val := range input {
		if !val {
			return false
		}
	}
	return true
}

func transposeBool(slice [][]bool) [][]bool {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]bool, xl)
	for i := range result {
		result[i] = make([]bool, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func ReduceMap(input map[int][]string) map[int]string {

	output := map[int]string{}
	for len(input) > 0 {

		// Generate & Sort Keys
		var keys []int
		for k := range input {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return len(input[keys[i]]) < len(input[keys[j]]) })

		if len(input[keys[0]]) != 1 {
			log.Fatal("Unable to determine next Rule.")
		}
		confirmedRule := input[keys[0]][0]
		output[keys[0]] = confirmedRule
		delete(input, keys[0])

		// Prune already found keys
		for _, k := range keys {
			index := -1
			for i, a := range input[k] {
				if a == confirmedRule {
					index = i
					break
				}
			}
			if index != -1 {
				input[k][index] = input[k][len(input[k])-1] // Copy last element to index i.
				input[k][len(input[k])-1] = ""              // Erase last element (write zero value).
				input[k] = input[k][:len(input[k])-1]       // Truncate slice.
			}

		}
	}
	return output
}
