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

func parsePlatform(input string) [][]rune {
	lines := strings.Split(input, "\n")
	platform := make([][]rune, len(lines))
	for x, line := range lines {
		platform[x] = make([]rune, len(line))
		for y, val := range line {
			platform[x][y] = val
		}
	}
	return platform
}

func tilt(platform [][]rune) [][]rune {
	tilted := make([][]rune, len(platform))
	for i := range platform {
		tilted[i] = make([]rune, len(platform[i]))
		copy(tilted[i], platform[i])
	}
	for y := 0; y < len(platform[0]); y++ {
		for x := 1; x < len(platform); x++ {
			if platform[x][y] != 'O' {
				continue
			}
			for i := x; i > 0 && tilted[i-1][y] == '.'; i-- {
				tilted[i][y], tilted[i-1][y] = tilted[i-1][y], tilted[i][y]
			}
		}
	}
	return tilted
}

func spin(platform [][]rune) [][]rune {
	spinned := make([][]rune, len(platform))
	for i := range platform {
		spinned[i] = make([]rune, len(platform[i]))
		copy(spinned[i], platform[i])
	}
	// North
	for y := 0; y < len(platform[0]); y++ {
		for x := 1; x < len(platform); x++ {
			if spinned[x][y] != 'O' {
				continue
			}
			for i := x; i > 0 && spinned[i-1][y] == '.'; i-- {
				spinned[i][y], spinned[i-1][y] = spinned[i-1][y], spinned[i][y]
			}
		}
	}
	// WEST
	for x := 0; x < len(platform); x++ {
		for y := 0; y < len(platform[0]); y++ {
			if spinned[x][y] != 'O' {
				continue
			}
			for j := y; j > 0 && spinned[x][j-1] == '.'; j-- {
				spinned[x][j], spinned[x][j-1] = spinned[x][j-1], spinned[x][j]
			}
		}
	}
	// SOUTH
	for y := 0; y < len(platform); y++ {
		for x := len(platform[0]) - 2; x >= 0; x-- {
			if spinned[x][y] != 'O' {
				continue
			}
			for i := x; i < len(platform)-1 && spinned[i+1][y] == '.'; i++ {
				spinned[i][y], spinned[i+1][y] = spinned[i+1][y], spinned[i][y]
			}
		}
	}
	// EAST
	for x := 0; x < len(platform); x++ {
		for y := len(platform[0]) - 2; y >= 0; y-- {
			if spinned[x][y] != 'O' {
				continue
			}
			for j := y; j < len(platform[0])-1 && spinned[x][j+1] == '.'; j++ {
				spinned[x][j], spinned[x][j+1] = spinned[x][j+1], spinned[x][j]
			}
		}
	}
	return spinned
}

func computeTotalLoad(platform [][]rune) (load int) {
	for x, line := range platform {
		for _, val := range line {
			if val != 'O' {
				continue
			}
			load += (len(platform) - x)
		}
	}
	return
}

func part1(input string) int {
	p := parsePlatform(input)
	t := tilt(p)
	for i := range p {
		fmt.Println(string(p[i]), " ", string(t[i]))
	}
	return computeTotalLoad(t)
}

func findCycle(input []int) (int, int) {
	for x := 0; x < len(input)/2; x++ {
		for k := 1; k < (len(input)-x)/2; k++ {
			found := true
			for i := 0; i < len(input)-x-k-1; i++ {
				if input[x+i%k] != input[x+i+k] {
					found = false
					break
				}
			}
			if found {
				return k, x
			}
		}
	}
	panic("no cycle found")
}

func part2(input string) int {
	p := parsePlatform(input)
	loads := make([]int, 4*len(p))
	for i := 0; i < 4*len(p); i++ {
		p = spin(p)
		loads[i] = computeTotalLoad(p)
	}
	cycleLength, offset := findCycle(loads)
	return loads[(1000000000-1-offset)%cycleLength+offset]
}

// 98793
