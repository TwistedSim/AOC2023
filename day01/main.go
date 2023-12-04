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

func findFirstDigit(str string) int {
    for _, char := range str {
        if char >= '0' && char <= '9' {
			return int(char - '0')
        }
    }
    return -1
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {
		r1 := findFirstDigit(line)
		r2 := findFirstDigit(util.Reverse(line))
		result += (10*r1 + r2)
	}
	return result
}

var r = strings.NewReplacer(
    "one",   "o1e",
    "two",   "t2o",
    "three", "t3e",
    "four",  "f4r",
    "five",  "f5e",
    "six",   "s6x",
    "seven", "s7n",
    "eight", "e8t",
    "nine",  "n9e",
)


func part2(input string) int {
	lines := strings.Fields(input)
	result := 0
	for _, line := range lines {
		patchedLine := r.Replace(line)
		patchedLine = r.Replace(patchedLine)
		r1 := findFirstDigit(patchedLine)
		r2 := findFirstDigit(util.Reverse(patchedLine))
		result += (10*r1 + r2)
	}
	return result
}