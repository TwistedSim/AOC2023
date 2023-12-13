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

func parsePattern(input string) [][]bool {
	lines := strings.Split(input, "\n")
	pattern := make([][]bool, len(lines))
	for i := range pattern {
		pattern[i] = make([]bool, len(lines[0]))
	}
	for x, line := range lines {
		for y, val := range line {
			if val == '#' {
				pattern[x][y] = true
			}
		}
	}
	return pattern
}

func compareLine(line1, line2 []bool) bool {
	for i := range line1 {
		if line1[i] != line2[i] {
			return false
		}
	}
	return true
}

func isValidReflection(pattern [][]bool, i int) bool {
	if i < 0 || i == len(pattern) - 1 {
		return false
	}
	for j := i+1; i >= 0 && j < len(pattern); i-- {
		if !compareLine(pattern[i], pattern[j]) {
			return false
		}
		j++
	}
	return true
}

func processPattern(pattern [][]bool) (points []int) {
	for i := range pattern {
		if isValidReflection(pattern, i) {
			points = append(points, (i+1)*100)
		}
	}
	pattern = util.Transpose(pattern)
	for i := range pattern {
		if isValidReflection(pattern, i) {
			points = append(points, i+1)
		}
	}
	return points
}

func part1(input string) (result int) {
	rawPatterns := strings.Split(input, "\n\n")
	for _, p := range rawPatterns {
		points := processPattern(parsePattern(p))
		result += points[0]
	}
	return result
}

func smudge(pattern [][]bool) int {
	rInit := processPattern(pattern)[0]
	for x := range pattern {
		for y := range pattern[0] {
			pattern[x][y] = !pattern[x][y]
			points := processPattern(pattern)
			pattern[x][y] = !pattern[x][y]
			if len(points) == 0 { continue }
			for _, p := range points {
				if p != rInit {
					return p
				}
			}
		}
	}
	panic("no smudge found")
}

func part2(input string) (result int) {
	rawPatterns := strings.Split(input, "\n\n")
	for _, p := range rawPatterns {
		result += smudge(parsePattern(p))
	}
	return result
}