/*
This is the solution to part 2,
I had a solution to part 1, but the technique used
didn't work at all for part 2. I scrapped it and started over.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Mask struct {
	Bits     []rune
	MemAddrs []MemAddr
}

type MemAddr struct {
	DecKey int
	BinKey []rune
	Value  int
}

func initMem(masks []Mask) map[int]int {
	out := make(map[int]int)
	for _, maskGroup := range masks {
		for _, memAddr := range maskGroup.MemAddrs {
			for i := range memAddr.BinKey {
				switch maskGroup.Bits[i] {
				case 'X':
					memAddr.BinKey[i] = 'X'
				case '1':
					memAddr.BinKey[i] = '1'
				case '0':
					continue
				}
			}
			var addrs []int
			permutateAddr(string(memAddr.BinKey), &addrs)
			for _, addr := range addrs {
				out[addr] = memAddr.Value
			}

		}

	}
	return out
}

func permutateAddr(addr string, out *[]int) {
	if strings.Index(addr, "X") == -1 {
		bin, err := strconv.ParseInt(addr, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		*out = append(*out, int(bin))

	} else {
		upAddr := strings.Replace(addr, "X", "1", 1)
		dnAddr := strings.Replace(addr, "X", "0", 1)
		permutateAddr(upAddr, out)
		permutateAddr(dnAddr, out)
	}
}

func sumMem(m map[int]int) int {
	var total int
	for _, v := range m {
		total += v
	}
	return total
}

func parseInput() []Mask {
	file, err := os.Open("input.prod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var masks []Mask
	for scanner.Scan() {
		text := scanner.Text()

		if text[:4] == "mask" {
			bits := []rune(strings.TrimPrefix(text, "mask = "))
			masks = append(masks, Mask{bits, []MemAddr{}})

		} else {
			text = strings.TrimPrefix(text, "mem[")
			line := strings.Split(text, "] = ")
			Deckey, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatal(err)
			}
			BinKey, err := digitToBin(Deckey)
			if err != nil {
				log.Fatal(err)
			}
			Value, err := strconv.Atoi(line[1])
			if err != nil {
				log.Fatal(err)
			}
			masks[len(masks)-1].MemAddrs = append(masks[len(masks)-1].MemAddrs, MemAddr{Deckey, BinKey, Value})

		}

	}
	return masks
}

func digitToBin(d int) ([]rune, error) {
	/*
		Takes an positive integer greater than zero and
		converts it to a splice of runes. Where the runes are
		`1` or `0`.

		The output splice has length == 36.
	*/
	length := 36
	var out = make([]rune, length)
	for i := range out {
		remainder := d - int(math.Pow(2, float64(length-(i+1))))
		if remainder >= 0 {
			out[i] = '1'
			d = remainder
		} else {
			out[i] = '0'
		}
	}
	if d != 0 {
		return nil, fmt.Errorf("Failed to completely convert %d", d)
	}
	return out, nil
}

func main() {
	ms := parseInput()
	mem := initMem(ms)
	fmt.Println(sumMem(mem))
}
