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

func expand(images [][]rune, expansionFactor int) (expandedImage [][]rune) {
	for _, line := range images {
		expandedImage = append(expandedImage, []rune(line))
		if !strings.Contains(string(line), "#") {
			for i := 0; i < expansionFactor-1; i++ {
				expandedImage = append(expandedImage, []rune(line))
			}
		}
	}
	return
}

func parseImage(input string, expansionFactor int) (galaxies [][]int) {
	image := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		image = append(image, []rune(line))
	}
	if expansionFactor > 1 {
		image = util.Transpose(expand(util.Transpose(expand(image, expansionFactor)), expansionFactor))
	}
	for x, lines := range image {
		for y, value := range lines {
			if value == '#' {
				galaxies = append(galaxies, []int{x, y})
			}
		}
	}
	return
}

func distances(galaxies [][]int) (result int) {
	for i := 0; i < len(galaxies); i++ {
		for j := i+1; j < len(galaxies); j++ {
			result += util.ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	return
}

func part1(input string) int {
	return distances(parseImage(input, 2))
}


func part2(input string) int {
	expandedDistance := distances(parseImage(input, 2))
	normalDistance := distances(parseImage(input, 1))
	expansionRate := 1000000
	return normalDistance + (expandedDistance - normalDistance) * (expansionRate - 1)
}