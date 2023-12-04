package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Grid struct {
	Rows int
	Cols int
	Data [][]rune
}

type Index struct {
	x int
	y int
}

type Number struct {
	index  Index
	value  int
	length int
}

func createGrid(input string) *Grid {
	lines := strings.Fields(input)
	rows := len(lines)
	cols := len(lines[0])
	data := make([][]rune, rows)
	for i, line := range lines {
		data[i] = make([]rune, cols)
		for j, val := range line {
			data[i][j] = rune(val)
		}
	}
	return &Grid{Rows: rows, Cols: cols, Data: data}
}

func (g *Grid) CheckNeighborsForSymbols(x int, y int, symbols string) bool {
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /* {x,y} */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < g.Rows && newY >= 0 && newY < g.Cols {
			if strings.ContainsRune(symbols, g.Data[newX][newY]) {
				return true
			}
		}
	}

	return false
}

func (g *Grid) FindSymbolsAroundNumber(number Number, symbols string) []Index {
	locs := []Index{}
	if number.index.y > 0 {
		if strings.ContainsRune(symbols, g.Data[number.index.x][number.index.y-1]) {
			locs = append(locs, Index{number.index.x, number.index.y - 1})
		}
	}
	if number.index.y+number.length < g.Cols-1 {
		if strings.ContainsRune(symbols, g.Data[number.index.x][number.index.y+number.length]) {
			locs = append(locs, Index{number.index.x, number.index.y + number.length})
		}
	}
	if number.index.x > 0 {
		for j := util.Max(number.index.y-1, 0); j <= util.Min(number.index.y+number.length, g.Cols-1); j++ {
			if strings.ContainsRune(symbols, g.Data[number.index.x-1][j]) {
				locs = append(locs, Index{number.index.x - 1, j})
			}
		}
	}
	if number.index.x < g.Rows-1 {
		for j := util.Max(number.index.y-1, 0); j <= util.Min(number.index.y+number.length, g.Cols-1); j++ {
			if strings.ContainsRune(symbols, g.Data[number.index.x+1][j]) {
				locs = append(locs, Index{number.index.x + 1, j})
			}
		}
	}
	return locs
}

/*
func (g *Grid) FindSymbolsAroundNumber(number Number, symbols string) []Index {
	directions := make([][]int, 6 + 2*number.length)
	directions[0] = []int{0, -1}
	directions[1] = []int{0, number.length}
	for j := -1; j <= number.length; j++ {
		directions[3+j] = []int{-1, j}
		directions[5+number.length+j] = []int{1, j}
	}

	locs := []Index{}
	for _, direction := range directions {
		x, y := number.index.x + direction[0], number.index.y + direction[1]
		if x < 0 || x >= g.Rows || y < 0 || y >= g.Cols {
			continue
		}
		if strings.ContainsRune(symbols, g.Data[x][y]) {
			locs = append(locs, Index{x, y})
		}
	}

	return locs
}
*/

func (g *Grid) PrintGrid() {
	for _, row := range g.Data {
		for _, val := range row {
			fmt.Printf("%c ", val)
		}
		fmt.Println()
	}
}

func (g *Grid) GenerateSymbolsMap(symbols string) map[Index][]Number {
	symbolsMap := map[Index][]Number{}
	for x, line := range g.Data {
		for _, numberBounds := range regexp.MustCompile(`\d+`).FindAllStringIndex(string(line), -1) {
			numberString := string(line)[numberBounds[0]:numberBounds[1]]
			numberValue, _ := strconv.Atoi(numberString)
			number := Number{
				index:  Index{x, numberBounds[0]},
				value:  numberValue,
				length: numberBounds[1] - numberBounds[0],
			}
			for _, index := range g.FindSymbolsAroundNumber(number, symbols) {
				if symbolsMap[index] == nil {
					symbolsMap[index] = []Number{}
				}
				symbolsMap[index] = append(symbolsMap[index], number)
			}
		}
	}
	return symbolsMap
}

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

func part1(input string) int {
	grid := createGrid(input)
	//grid.PrintGrid()
	symbolsMap := grid.GenerateSymbolsMap("*#$+@/%&-=")
	result := 0
	for _, numbers := range symbolsMap {
		for _, number := range numbers {
			result += number.value
		}
	}

	return result
}

func part2(input string) int {
	grid := createGrid(input)
	//grid.PrintGrid()
	symbolsMap := grid.GenerateSymbolsMap("*")
	result := 0
	for _, numbers := range symbolsMap {
		if len(numbers) == 2 {
			result += numbers[0].value * numbers[1].value
		}
	}

	return result
}

func part1alt(input string) int {
	result := 0
	grid := createGrid(input)
	//grid.PrintGrid()
	for x, line := range grid.Data {
		for _, numberBounds := range regexp.MustCompile(`\d+`).FindAllStringIndex(string(line), -1) {
			for y := numberBounds[0]; y < numberBounds[1]; y++ {
				if grid.CheckNeighborsForSymbols(x, y, "*#$+@/%&-=") {
					number, _ := strconv.Atoi(string(line)[numberBounds[0]:numberBounds[1]])
					result += number
					break
				}
			}
		}
	}
	return result
}

func part2alt(input string) int {
	grid := createGrid(input)
	result := 0
	for x, line := range grid.Data {
		for y, val := range line {
			if val != '*' {
				continue
			}
			grid.Data[x][y] = 'T'
			numbers := []int{}
			for i, line := range grid.Data[util.Max(0, x-1) : util.Min(grid.Rows-1, x+1)+1] {
				for _, numberBounds := range regexp.MustCompile(`\d+`).FindAllStringIndex(string(line), -1) {
					for y := numberBounds[0]; y < numberBounds[1]; y++ {
						if grid.CheckNeighborsForSymbols(x+i-1, y, "T") {
							numberString := string(line)[numberBounds[0]:numberBounds[1]]
							number, _ := strconv.Atoi(numberString)
							numbers = append(numbers, number)
							break
						}
					}
				}
			}
			grid.Data[x][y] = '*'
			if len(numbers) == 2 {
				result += numbers[0] * numbers[1]
			}
		}
	}
	return result
}
