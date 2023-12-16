package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"strconv"
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

func hash(input string) (result int) {
	for _, c := range input {
		result += int(c)
		result *= 17
		result %= 256
	}
	return
}

func part1(input string) (result int) {
	for _, step := range strings.Split(input, ",") {
		result += hash(step)
	}
	return
}

type Lense struct {
	Label       string
	FocalLength int
}

func parseLense(input string) Lense {
	parts := strings.Split(input, "=")
	focalLength, _ := strconv.Atoi(parts[1])
	return Lense{
		Label:       parts[0],
		FocalLength: focalLength,
	}
}

func contains(lenses []Lense, label string) int {
	for idx, l := range lenses {
		if l.Label == label {
			return idx
		}
	}
	return -1
}

func part2(input string) (result int) {
	boxes := map[int][]Lense{}
	for _, step := range strings.Split(input, ",") {
		if strings.Contains(step, "=") {
			lense := parseLense(step)
			h := hash(lense.Label)
			if idx := contains(boxes[h], lense.Label); idx != -1 {
				boxes[h][idx] = lense
			} else {
				boxes[h] = append(boxes[h], lense)
			}
		} else { // ends with -
			label := step[:len(step)-1]
			h := hash(label)
			if idx := contains(boxes[h], label); idx != -1 {
				boxes[h] = append(boxes[h][:idx], boxes[h][idx+1:]...)
			}
		}
	}
	for boxIdx, lenses := range boxes {
		for lenseIdx, lense := range lenses {
			result += (boxIdx + 1) * (lenseIdx + 1) * lense.FocalLength
		}
	}
	return
}
