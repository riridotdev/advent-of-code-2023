package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	totalScore := 0

	cards := strings.Split(inputString, "\n")
	for _, card := range cards {
		parts := strings.Split(card, "|")

		winningNumbersList := strings.Fields(strings.TrimSpace(strings.Split(parts[0], ":")[1]))

		winningNumbers := map[string]struct{}{}
		for _, winningNumber := range winningNumbersList {
			winningNumbers[winningNumber] = struct{}{}
		}

		score := 0

		numbersList := strings.Fields(strings.TrimSpace(parts[1]))
		for _, number := range numbersList {
			if _, ok := winningNumbers[number]; !ok {
				continue
			}
			score += max(score, 1)
		}

		totalScore += score
	}

	fmt.Println(totalScore)
}
