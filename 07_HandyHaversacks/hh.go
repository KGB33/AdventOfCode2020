package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	g := parseInput()
	seenBags := make(map[string]bool)
	countOuterBags("shiny gold bag", g, seenBags)
	fmt.Println("Part 1: ", len(seenBags))
	fmt.Println("Part 2: ", countInnerBags("shiny gold bag", g)-1)
}

type Graph struct {
	bags  map[string]bool // Set of bags
	edges []Edge
}

func NewGraph() Graph {
	return Graph{make(map[string]bool), *new([]Edge)}
}

func (g Graph) addBag(b string) {
	g.bags[b] = true

}

func (g *Graph) addEdge(e Edge) {
	g.edges = append(g.edges, e)
}

func (g Graph) String() string {
	output := ""
	for _, e := range g.edges {
		output += e.String()
	}
	return output
}

type Edge struct {
	weight int
	start  string //bag
	end    string // bag
}

func (e Edge) String() string {
	return fmt.Sprintf("%d :: %s ---> %s\n", e.weight, e.start, e.end)
}

func countOuterBags(b string, g Graph, seen map[string]bool) {
	for _, e := range g.edges {
		if e.end == b && !seen[e.start] {
			seen[e.start] = true
			countOuterBags(e.start, g, seen)
		}
	}
}

func countInnerBags(b string, g Graph) int {
	/*
		This function returns the number of bags within
		a given bag, including the bag itself.
	*/
	total := 1
	for _, e := range g.edges {
		if e.start == b {
			total += countInnerBags(e.end, g) * e.weight
		}
	}
	return total
}

func parseInput() Graph {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := NewGraph()
	rgx := regexp.MustCompile(` contain |, `)
	for scanner.Scan() {
		text := scanner.Text()
		text = text[:len(text)-1]
		text = strings.Replace(text, "bags", "bag", -1)
		bags := rgx.Split(text, -1)
		parent := bags[0]
		graph.addBag(parent)
		for _, child := range bags[1:] {
			if child == "no other bag" {
				continue
			}
			amount := int(child[0] - '0')
			graph.addEdge(Edge{amount, parent, child[2:]})
		}
	}

	return graph
}
