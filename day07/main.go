package main

import (
	"aoc/util"
	"flag"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
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

const (
	fiveOfAKind  = 6
	fourOfAKind  = 5
	fullHouse    = 4
	threeOfAKind = 3
	twoPair      = 2
	onePair      = 1
	highCard     = 0
)

type Game struct {
	Cards string
	Bid   int
	Type  int
}

var mapCards = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func histogram(cards string) map[rune]int {
	hist := make(map[rune]int, len(cards))
	for _, card := range cards {
		if _, ok := hist[card]; !ok {
			hist[card] = 0
		}
		hist[card] += 1
	}
	return hist
}

func getCardType(card string, joker bool) int {
	if joker {
		jokerCount := strings.Count(card, "J")
		card = strings.ReplaceAll(card, "J", "")
		return getCardTypeNormal(card + getBestJoker(card, jokerCount))
	}
	return getCardTypeNormal(card)
}

func getCardTypeNormal(card string) int {
	hist := histogram(card)
	values := maps.Values(hist)
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	switch values[0] {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if values[1] == 2 {
			return fullHouse
		}
		return threeOfAKind
	case 2:
		if values[1] == 2 {
			return twoPair
		}
		return onePair
	}
	return highCard
}

func getBestJoker(card string, jokerCount int) string {
	hist := histogram(card)
	maxValue := 0
	var bestCard rune
	for key, value := range hist {
		if value > maxValue {
			maxValue = value
			bestCard = key
		}
	}
	return strings.Repeat(string(bestCard), jokerCount)
}

func parseGames(input string, joker bool) (games []Game) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		games = append(games, Game{
			Cards: parts[0],
			Bid:   util.ParseNumbers(parts[1])[0],
			Type:  getCardType(parts[0], joker),
		})
	}
	return
}

func compareCards(g1, g2 Game, joker bool) bool {
	for i := 0; i < len(g1.Cards); i++ {
		c1, c2 := mapCards[rune(g1.Cards[i])], mapCards[rune(g2.Cards[i])]
		if c1 == c2 {
			continue
		}
		if joker {
			if c1 == mapCards['J'] {
				return true
			}
			if c2 == mapCards['J'] {
				return false
			}
		}
		return c1 < c2
	}
	return false
}

func rankGames(games []Game, joker bool) []Game {
	sort.SliceStable(games, func(i, j int) bool {
		return compareCards(games[i], games[j], joker)
	})
	sort.SliceStable(games, func(i, j int) bool {
		return games[i].Type < games[j].Type
	})
	return games
}

func part1(input string) (result int) {
	games := parseGames(input, false)
	games = rankGames(games, false)
	for rank, game := range games {
		result += game.Bid * (rank + 1)
	}
	return
}

func part2(input string) (result int) {
	games := parseGames(input, true)
	games = rankGames(games, true)
	for rank, game := range games {
		result += game.Bid * (rank + 1)
	}
	return
}
