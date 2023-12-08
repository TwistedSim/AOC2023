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

func (g *Graph) AddVertex(v string) {
	g.Vertices[v] = &Vertex{Value: v}
}

func (g *Graph) AddEdges(srcKey, leftKey, rightKey string) {
	g.Vertices[srcKey].Left = g.Vertices[leftKey]
	g.Vertices[srcKey].Right = g.Vertices[rightKey]
}

func makeGraph(list map[string][]string) *Graph {
	g := &Graph{Vertices: map[string]*Vertex{}}
	for vertex, edges := range list {
		if _, ok := g.Vertices[vertex]; !ok {
			g.AddVertex(vertex)
		}
		for _, edge := range edges {
			if _, ok := g.Vertices[edge]; !ok {
				g.AddVertex(edge)
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

func part1(input string) int {
	graph, directions := parseMaps(input)
	vertex, i := graph.Vertices["AAA"], 0
	for i = 0; vertex.Value != "ZZZ"; i++ {
		switch directions[i%len(directions)] {
			case 'L': vertex = vertex.Left
			case 'R': vertex = vertex.Right
		}
	}
	return i
}

func part2(input string) int {
	graph, directions := parseMaps(input)
	result, i := 1, 1
	for _, v := range graph.Vertices {
		if v.Value[2] != 'A' { continue }
		for i = 0; v.Value[2] != 'Z'; i++ {
			switch directions[i%len(directions)] {
				case 'L': v = v.Left
				case 'R': v = v.Right
			}
		}
		result = util.LCM(result, i)
	}
	return result
}
