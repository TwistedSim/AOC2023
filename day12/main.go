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

type Record struct {
	Springs []rune // #, ? , .
	Groups  []int
}

func parseRecords(input string, folded bool) (records []Record) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		springs := parts[0]
		groups := parts[1]
		if !folded {
			factor := 5
			unfoldedSprings := make([]string, factor)
			unfoldedGroups := make([]string, factor)
			for i := 0; i < factor; i++ {
				unfoldedSprings[i] = springs
				unfoldedGroups[i] = groups
			}
			springs = strings.Join(unfoldedSprings, "?")
			groups = strings.Join(unfoldedGroups, ",")
		}
		records = append(records, Record{
			Springs: []rune(springs),
			Groups:  util.ParseNumbers(strings.ReplaceAll(groups, ",", " ")),
		})
	}
	return
}

func processRecord(record Record) (result int) {
	springs := []bool{false}
	for _, g := range record.Groups {
		for i := 0; i < g; i++ {
			springs = append(springs, true)
		}
		springs = append(springs, false)
	}
	chars := append([]rune{'.'}, record.Springs...)
	chars = append(chars, '.')

	n := len(chars)
	m := len(springs)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	dp[n][m] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			damaged, operational := false, false
			switch chars[i] {
			case '#':
				damaged = true
			case '.':
				operational = true
			case '?':
				operational, damaged = true, true
			}
			sum := 0
			if damaged && springs[j] {
				sum += dp[i+1][j+1]
			} else if operational && !springs[j] {
				sum += dp[i+1][j+1] + dp[i+1][j]
			}
			dp[i][j] = sum
		}
	}
	for _, line := range dp {
		for _, val := range line {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}

	return dp[0][0]
}

func part1(input string) (result int) {
	records := parseRecords(input, true)
	for _, record := range records {
		result += processRecord(record)
	}
	return
}

func part2(input string) (result int) {
	records := parseRecords(input, false)
	for _, record := range records {
		result += processRecord(record)
	}
	return
}
