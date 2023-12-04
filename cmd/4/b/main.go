package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	cards := strings.Split(inputString, "\n")

	queue := []int{}
	for i := range cards {
		queue = append(queue, i)
	}

	totalCards := 0
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		card := cards[current]
		parts := strings.Split(card, "|")
		winningNumbersList := strings.Fields(strings.TrimSpace(strings.Split(parts[0], ":")[1]))

		winningNumbers := map[string]struct{}{}
		for _, winningNumber := range winningNumbersList {
			winningNumbers[winningNumber] = struct{}{}
		}

		winningNumberCount := 0
		numbersList := strings.Fields(strings.TrimSpace(parts[1]))
		for _, number := range numbersList {
			if _, ok := winningNumbers[number]; !ok {
				continue
			}
			winningNumberCount += 1
		}

		for i := 1; i <= winningNumberCount; i++ {
			queue = append(queue, current+i)
		}

		totalCards += 1
	}

	fmt.Println(totalCards)
}
