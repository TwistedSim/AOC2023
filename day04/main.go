package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	Numbers        []int
	WinningNumbers []int
}

func (c *Card) matchingNumbers() (matches int) {
	return len(util.Intersection(c.Numbers, c.WinningNumbers))
}

func parseNumbers(numbersString string) (numbers []int) {
	for _, numberString := range strings.Fields(numbersString) {
		number, _ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}
	return numbers
}

func parseCards(input string) (cards []Card) {
	for _, cardString := range strings.Split(input, "\n") {
		cardParts := strings.Split(cardString, ":")
		id, _ := strconv.Atoi(strings.Fields(cardParts[0])[1])
		numbersParts := strings.Split(cardParts[1], "|")
		cards = append(cards, Card{
			Id:             id,
			Numbers:        parseNumbers(numbersParts[0]),
			WinningNumbers: parseNumbers(numbersParts[1]),
		})
	}
	return cards
}

func computeCardRewards(card Card) []int {
	reward := make([]int, card.matchingNumbers())
	for i := range reward {
		reward[i] = i + 1 + card.Id
	}
	return reward
}

func computeWinningCards(cards []Card) (winningCards int) {
	totalRewards := make([]int, len(cards))
	for i := len(cards) - 1; i >= 0; i-- {
		cardRewards := computeCardRewards(cards[i])
		totalRewards[i] = len(cardRewards)
		for _, reward := range cardRewards {
			totalRewards[i] += totalRewards[reward-1]
		}
		winningCards += totalRewards[i]
	}
	return
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

func part1(input string) (result int) {
	cards := parseCards(input)
	for _, card := range cards {
		result += int(math.Pow(2, float64(card.matchingNumbers()-1)))
	}
	return
}

func part2(input string) (result int) {
	cards := parseCards(input)
	rewards := computeWinningCards(cards)
	return rewards + len(cards)
}
