package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

type game struct {
	id       int
	maxRed   int
	maxBlue  int
	maxGreen int
}

func parseGame(str string) game {
	res := game{}
	fmt.Sscanf(str, "Game %d", &res.id)
	data := strings.Split(str, ":")[1]
	for _, s := range regexp.MustCompile(`[,;]+`).Split(data, -1) {
		var (
			qty int
			color string
		)
		fmt.Sscanf(s, "%d %s", &qty, &color)
		switch color {
		case "red":
			res.maxRed = util.Max(res.maxRed, qty)
		case "blue":
			res.maxBlue = util.Max(res.maxBlue, qty)
		case "green":
			res.maxGreen = util.Max(res.maxGreen, qty)
		}
	}
	return res
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	result := 0
	for _,gameStr := range strings.Split(input, "\n") {
		game := parseGame(gameStr)
		if game.maxRed > 12 || game.maxBlue > 14 || game.maxGreen > 13 {
			continue
		}
		result += game.id
	}

	return result
}


func part2(input string) int {
	result := 0
	for _,gameStr := range strings.Split(input, "\n") {
		game := parseGame(gameStr)
		result += game.maxRed * game.maxBlue * game.maxGreen
	}

	return result
}