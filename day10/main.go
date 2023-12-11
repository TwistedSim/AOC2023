package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"strings"
)

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

func parseMap(input string) ([][]string, []int) {
	lines := strings.Split(input, "\n")
	pipes := make([][]string, len(lines))
	startPos := []int{0, 0}
	for i, line := range lines {
		pipes[i] = make([]string, len(line))
		for j, val := range line {
			if val == 'S' {
				startPos = []int{i, j}
			}
			pipes[i][j] = string(val)

		}
	}
	return pipes, startPos
}

var directions = map[string][]int{
	"N": {-1, 0},
	"S": {1, 0},
	"E": {0, 1},
	"W": {0, -1},
}

var connectionMap = map[string]map[string]string{
	"N": {"|": "N", "7": "W", "F": "E"},
	"S": {"|": "S", "L": "E", "J": "W"},
	"E": {"-": "E", "7": "S", "J": "N"},
	"W": {"-": "W", "L": "N", "F": "S"},
}

func findInitalDirection(pipes [][]string, startPos []int) (string, []int) {
	for _, cardinalDirection := range []string{"N", "E", "S", "W"} {
		xyDirection := directions[cardinalDirection]
		x, y := startPos[0]+xyDirection[0], startPos[1]+xyDirection[1]
		if x < 0 || y < 0 || x >= len(pipes[0]) || y >= len(pipes[0]) {
			continue
		}
		pipe := string(pipes[x][y])
		startCardinalDirection, ok := connectionMap[cardinalDirection][pipe]
		if ok {
			return startCardinalDirection, []int{x, y}
		}
	}
	return "", []int{}
}

func part1(input string) int {
	pipes, startPos := parseMap(input)
	cardinalDirection, pos := findInitalDirection(pipes, startPos)

	cycleLength := 0
	for pipe := ""; pipe != "S"; {
		cycleLength++
		xyDirection := directions[cardinalDirection]
		pos[0], pos[1] = pos[0]+xyDirection[0], pos[1]+xyDirection[1]
		pipe = string(pipes[pos[0]][pos[1]])
		cardinalDirection = connectionMap[cardinalDirection][pipe]
	}

	return int(cycleLength/2) + 1
}

func part2(input string) int {
	pipes, startPos := parseMap(input)
	cardinalDirection, pos := findInitalDirection(pipes, startPos)

	isPath := make([][]bool, len(pipes))
	for i := range pipes {
		isPath[i] = make([]bool, len(pipes[0]))
	}

	isPath[startPos[0]][startPos[1]] = true
	isPath[pos[0]][pos[1]] = true
	for pipe := ""; pipe != "S"; {
		xyDirection := directions[cardinalDirection]
		pos[0], pos[1] = pos[0]+xyDirection[0], pos[1]+xyDirection[1]
		pipe = string(pipes[pos[0]][pos[1]])
		cardinalDirection = connectionMap[cardinalDirection][pipe]
		isPath[pos[0]][pos[1]] = true
	}

	containedTiles := 0
	for x, lines := range pipes {
		for y := range lines {
			if isPath[x][y] {
				continue
			}
			// check in a diagonal if we have a even or odd amount of crossing
			x2, y2 := x, y
			crossing := 0
			for x2 < len(pipes) && y2 < len(pipes[0]) {
				tile := pipes[x2][y2]
				if isPath[x2][y2] && tile != "L" && tile != "7" {
					crossing++
				}
				x2++
				y2++
			}
			if crossing%2 == 1 {
				containedTiles++
			}
		}
	}

	return containedTiles
}
