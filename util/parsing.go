package util

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseStringFromRegex(regEx *regexp.Regexp, content string) map[string]string {
    match := regEx.FindStringSubmatch(content)
    paramsMap := make(map[string]string)
    for i, name := range regEx.SubexpNames() {
        if i > 0 && i <= len(match) {
            paramsMap[name] = match[i]
        }
    }
    return paramsMap
}

func ParseNumbers(numbersString string) (numbers []int) {
	for _, numberString := range strings.Fields(numbersString) {
		number, _ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}
	return numbers
}

func ParseRuneGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}
