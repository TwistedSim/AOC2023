package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	Value string
	Left  *Vertex
	Right *Vertex
}

func (g *Graph) AddVertex(key, val string) {
	g.Vertices[key] = &Vertex{Value: val}
}

func (g *Graph) AddEdges(srcKey, leftKey, rightKey string) {
	g.Vertices[srcKey].Left = g.Vertices[leftKey]
	g.Vertices[srcKey].Right = g.Vertices[rightKey]

}

func makeGraph(list map[string][]string) *Graph {
	g := &Graph{Vertices: map[string]*Vertex{}}
	for vertex, edges := range list {
		if _, ok := g.Vertices[vertex]; !ok {
			g.AddVertex(vertex, vertex)
		}
		for _, edge := range edges {
			if _, ok := g.Vertices[edge]; !ok {
				g.AddVertex(edge, edge)
			}
		}
		g.AddEdges(vertex, edges[0], edges[1])
	}
	return g
}

func parseMaps(input string) (graph *Graph, directions string) {
	parts := strings.Split(input, "\n\n")
	direction := parts[0]
	lines := strings.Split(parts[1], "\n")
	adjencyMatrix := make(map[string][]string, len(lines))
	mapRegex := regexp.MustCompile(`[A-Z0-9]{3}`)
	for _, line := range lines {
		val := mapRegex.FindAllString(line, -1)
		adjencyMatrix[val[0]] = val[1:]
	}
	return makeGraph(adjencyMatrix), direction
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) (result int) {
	graph, directions := parseMaps(input)
	vertex := graph.Vertices["AAA"]
	for vertex.Value != "ZZZ" {
		if directions[result%len(directions)] == 'L' {
			vertex = vertex.Left
		} else {
			vertex = vertex.Right
		}
		result++
	}
	return result
}

func part2(input string) (result int) {
	graph, directions := parseMaps(input)
	results := []int{}
	for _, v := range graph.Vertices {
		if v.Value[2] == 'A' {
			dirIndex := 0
			for v.Value[2] != 'Z' {
				if directions[dirIndex%len(directions)] == 'L' {
					v = v.Left 
				} else {
					v = v.Right 
				}
				dirIndex++
			}
			results = append(results, dirIndex)
		}
	}
	return util.LCM(results...)
}
