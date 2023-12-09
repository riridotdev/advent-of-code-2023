package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	rank  handRank
	bid   int
}

type handRank int

const (
	_ handRank = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	handStrings := strings.Split(inputString, "\n")

	hands := []hand{}
	for _, handString := range handStrings {
		handParts := strings.Fields(handString)
		cards := handParts[0]
		handRank := rankHand(cards)
		bid, _ := strconv.Atoi(handParts[1])
		hands = append(hands, hand{cards, handRank, bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank > hands[j].rank {
			return false
		}
		if hands[i].rank < hands[j].rank {
			return true
		}

		iCards := hands[i].cards
		jCards := hands[j].cards
		for cardIdx := range iCards {
			if cardValue(iCards[cardIdx]) > cardValue(jCards[cardIdx]) {
				return false
			}
			if cardValue(iCards[cardIdx]) < cardValue(jCards[cardIdx]) {
				return true
			}
		}
		return false
	})

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.bid * (1 + i)
	}

	fmt.Println(totalWinnings)
}

func rankHand(hand string) handRank {
	cards := map[rune]int{}
	wildCards := 0
	for _, card := range hand {
		if card == 'J' {
			wildCards += 1
			continue
		}
		cards[card] += 1
	}

	cardCounts := []int{}
	for _, cardCount := range cards {
		cardCounts = append(cardCounts, cardCount)
	}

	switch len(cardCounts) {
	case 0:
		return FiveOfAKind
	case 1:
		return FiveOfAKind
	case 2:
		for _, count := range cardCounts {
			if count+wildCards == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for _, count := range cardCounts {
			if count+wildCards == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func cardValue(r byte) int {
	switch r {
	case 'T':
		return 10
	case 'J':
		return 0
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return int(r - '0')
	}
}
