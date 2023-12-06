package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"math"
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

type Race struct {
	BestDistance int
	MaxTime      int
}

func parseRaces(input string) (races []Race) {
	lines := strings.Split(input, "\n")
	distances := util.ParseNumbers(strings.Split(lines[1], ":")[1])
	times := util.ParseNumbers(strings.Split(lines[0], ":")[1])
	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			BestDistance: distances[i],
			MaxTime:      times[i],
		})
	}
	return
}

func parseBigRace(input string) (race Race) {
	lines := strings.Split(input, "\n")
	distanceStr := strings.Split(lines[1], ":")[1]
	timeStr := strings.Split(lines[0], ":")[1]
	distance := strings.ReplaceAll(distanceStr, " ", "")
	time := strings.ReplaceAll(timeStr, " ", "")
	race.BestDistance, _ = strconv.Atoi(distance)
	race.MaxTime, _ = strconv.Atoi(time)
	return
}

func quad(Di, Ti int) (roots []float64) {
	T := float64(Ti)
	D := float64(Di)
	det := math.Sqrt(math.Pow(T, 2) - 4*D)
	return []float64{
		(T - det) / 2,
		(T + det) / 2,
	}
}

func computeNumberOfWinningRaces(race Race) (wins int) {
	xs := quad(race.BestDistance, race.MaxTime)
	return int(math.Ceil(xs[1]-1) - math.Floor(xs[0]+1) + 1)
}

func part1(input string) int {
	races := parseRaces(input)
	result := 1
	for _, race := range races {
		result *= computeNumberOfWinningRaces(race)
	}
	return result
}

func part2(input string) int {
	race := parseBigRace(input)
	return computeNumberOfWinningRaces(race)
}
