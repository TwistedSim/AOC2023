package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"math"
	"regexp"
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

func parseCards(input string) (cards []Card) {
	cardRegex := regexp.MustCompile(`Card\s+(?P<Id>\d+):(?P<Numbers>(\s+\d+)+)\s\|(?P<WinningNumbers>(\s+\d+)+)+`)
	for _, cardString := range strings.Split(input, "\n") {
		params := util.ParseStringFromRegex(cardRegex, cardString)
		cards = append(cards, Card{
			Id:             util.ParseNumbers(params["Id"])[0],
			Numbers:        util.ParseNumbers(params["Numbers"]),
			WinningNumbers: util.ParseNumbers(params["WinningNumbers"]),
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
	result = computeWinningCards(cards) + len(cards)
	return
}
