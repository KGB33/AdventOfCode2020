package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	batch := parseInput()

	// Part 1
	p1 := trimInvalid(batch)
	fmt.Printf("Part 1: %d\n", len(p1))

	// Part 2
	passports := parsePassports(p1)
	p2 := part2(passports)
	fmt.Printf("Part 2: %d\n", p2)
}

func parseInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := []string{""}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			result = append(result, "")
			continue
		}
		result[len(result)-1] += (text + " ")
	}
	return result
}

func part2(passes []passport) int {
	var count int
	for _, pass := range passes {
		if pass.vaidate() {
			count++
		}
	}
	return count
}

func trimInvalid(data []string) []string {
	var output []string
	for _, val := range data {
		if containsRequiredKeys(val) {
			output = append(output, val)
		}
	}
	return output
}

func parsePassports(data []string) []passport {
	output := make([]passport, len(data))

	for i, p := range data {
		pairs := strings.Fields(p)
		var pass passport
		for _, pair := range pairs {
			kv := strings.Split(pair, ":")
			key, value := kv[0], kv[1]

			switch key {
			case "cid":
				continue
			case "byr":
				byr, err := strconv.Atoi(value)
				if err != nil {
					log.Fatal(err)
				}
				pass.byr = byr

			case "iyr":
				iyr, err := strconv.Atoi(value)
				if err != nil {
					log.Fatal(err)
				}
				pass.iyr = iyr
			case "eyr":
				eyr, err := strconv.Atoi(value)
				if err != nil {
					log.Fatal(err)
				}
				pass.eyr = eyr

			case "hgt":
				if strings.Contains(value, "in") {
					hgt, err := strconv.Atoi(strings.TrimSuffix(value, "in"))
					if err != nil {
						log.Fatal(err)
					}
					pass.hgt = hgt
					pass.hgtUnit = "in"
				} else if strings.Contains(value, "cm") {
					hgt, err := strconv.Atoi(strings.TrimSuffix(value, "cm"))
					if err != nil {
						log.Fatal(err)
					}
					pass.hgt = hgt
					pass.hgtUnit = "cm"
				} else {
					pass.hgt = -1
					pass.hgtUnit = "in"
				}
			case "ecl":
				pass.ecl = value
			case "hcl":
				pass.hcl = value
			case "pid":
				pass.pid = value
			default:
				log.Fatalf("No case for %s", key)
			}

		}
		output[i] = pass

	}

	return output
}

type passport struct {
	byr int
	iyr int
	eyr int

	hgt     int
	hgtUnit string

	ecl string
	hcl string

	pid string
}

func (p passport) vaidate() bool {
	//  Year stuff
	if (p.byr < 1920) || (p.byr > 2002) {
		return false
	}
	if (p.iyr < 2010) || (p.iyr > 2020) {
		return false
	}
	if (p.eyr < 2020) || (p.eyr > 2030) {
		return false
	}

	// Height
	switch p.hgtUnit {
	case "cm":
		if (p.hgt < 150) || (p.hgt > 193) {
			return false
		}
	case "in":
		if (p.hgt < 59) || (p.hgt > 76) {
			return false
		}
	default:
		return false
	}

	// Hair Color
	if len(p.hcl) != 7 {
		return false
	}
	matches, err := regexp.MatchString("^#[a-f0-9]+$", p.hcl)
	if err != nil {
		log.Fatal(err)
	}
	if !matches {
		return false
	}

	// Eye Color
	switch p.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		break
	default:
		return false
	}

	// Passport id
	if len(p.pid) != 9 {
		return false
	}
	if _, err := strconv.Atoi(p.pid); err != nil {
		return false
	}

	// Default case
	return true
}

func containsRequiredKeys(passport string) bool {
	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid",
	}

	for _, sub := range required {
		if !strings.Contains(passport, sub) {
			return false
		}
	}
	return true
}
