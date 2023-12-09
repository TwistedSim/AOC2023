package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"math"
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

func parseHistories(input string) [][]int {
	histories := strings.Split(input, "\n")
	dataset := make([][]int, len(histories))
	for i, history := range strings.Split(input, "\n") {
		dataset[i] = util.ParseNumbers(history)
	}
	return dataset
}

func diff(list []int) []int {
	res := make([]int, len(list)-1)
	for i := 0; i<len(list)-1; i++ {
		res[i] = list[i+1] - list[i]
	}
	return res
}

func transposeSum(list [][]int) []int {
	tSum := make([]int, len(list[0]))
	listT := util.Transpose(list)
	for i := 0; i < len(listT); i++ {
		tSum[i] = util.Sum(listT[i])
	}
	return tSum
}

func part1(input string) (result int) {
	dataset := parseHistories(input)
	datasetSum := transposeSum(dataset)
	dataSetDiff := datasetSum
	for i := 0; util.Sum(dataSetDiff) != 0; i++ {
		result += dataSetDiff[len(dataSetDiff)-1]
		dataSetDiff = diff(dataSetDiff)
	}
	return
}


func part2(input string) (result int) {
	dataset := parseHistories(input)
	datasetSum := transposeSum(dataset)
	dataSetDiff := datasetSum
	for i := 0; util.Sum(dataSetDiff) != 0; i++ {
		result += int(math.Pow(-1, float64(i))) * dataSetDiff[0]
		dataSetDiff = diff(dataSetDiff)
	}
	return
}