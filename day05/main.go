package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"sort"
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

func createAlmanac(parts []string) []IntervalTree {
	almanac := make([]IntervalTree, len(parts))
	for partIdx, part := range parts {
		lines := strings.Split(part, "\n")
		for _, mapStr := range lines[1:] {
			info := util.ParseNumbers(mapStr)
			almanac[partIdx].insert(Interval{
				Low:    info[1],
				High:   info[1] + info[2],
				Offset: info[0] - info[1],
			})
		}
	}
	return almanac
}

func part1(input string) (result int) {
	parts := strings.Split(input, "\n\n")
	seeds := util.ParseNumbers(strings.Split(parts[0], "seeds:")[1])
	almanac := createAlmanac(parts[1:])
	locations := seeds
	for _, page := range almanac {
		for i, location := range locations {
			interval := page.search(location)
			locations[i] = location + interval.Offset
		}
	}
	result, _ = util.MinMax(locations)
	return
}

func part2(input string) (result int) {
	parts := strings.Split(input, "\n\n")
	seedsInfo := util.ParseNumbers(strings.Split(parts[0], "seeds:")[1])
	seeds := make([]*Interval, len(seedsInfo)/2)
	for i, j := 0, 0; i < len(seedsInfo); i += 2 {
		seeds[j] = &Interval{
			Low:  seedsInfo[i],
			High: seedsInfo[i] + seedsInfo[i+1],
		}
		j++
	}

	almanac := createAlmanac(parts[1:])

	currentIntervals := seeds
	for _, page := range almanac {
		nextIntervals := []*Interval{}
		for _, interval := range currentIntervals {
			matchingIntervals := page.searchInterval(interval)

			if len(matchingIntervals) == 0 {
				nextIntervals = append(nextIntervals, interval)
				continue
			}

			sort.SliceStable(matchingIntervals, func(i, j int) bool {
				return matchingIntervals[i].Low < matchingIntervals[j].Low
			})

			lowestMatch := matchingIntervals[0].Low
			if interval.Low < lowestMatch {
				nextIntervals = append(nextIntervals, &Interval{Low: interval.Low, High: lowestMatch})
			}
			
			
			largestMatch := matchingIntervals[len(matchingIntervals)-1].High
			if interval.High > largestMatch {
				nextIntervals = append(nextIntervals, &Interval{Low:  largestMatch, High: interval.High})
			}
			
			for _, m := range matchingIntervals {
				nextIntervals = append(nextIntervals, &Interval{
					Low:  util.Max(interval.Low, m.Low) + m.Offset,
					High: util.Min(interval.High, m.High) + m.Offset,
				})
			}
			
			for i := 0; i < len(matchingIntervals)-1; i++ {
				m1, m2 := matchingIntervals[i], matchingIntervals[i+1]
				if m2.Low > m1.High {
					nextIntervals = append(nextIntervals, &Interval{Low: m1.High, High: m2.Low})
				}
			}
			
		}
		currentIntervals = nextIntervals
	}

	for _, i := range currentIntervals {fmt.Println(i) }
	result = currentIntervals[0].Low
	for _, seed := range currentIntervals {
		if seed.Low < result {
			result = seed.Low
		}
	}
	return
}
