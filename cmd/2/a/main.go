package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var counts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	games := strings.Split(inputString, "\n")

	validGameIdxSum := 0
gameLoop:
	for gameIdx, game := range games {
		game = strings.SplitAfter(game, ":")[1]
		combinations := strings.Split(game, ";")

		for _, combination := range combinations {
			colourCounts := strings.Split(combination, ",")
			for _, colourCount := range colourCounts {
				parts := strings.Split(strings.TrimSpace(colourCount), " ")

				countString := parts[0]
				count, _ := strconv.Atoi(countString)

				colour := parts[1]

				if counts[colour] < count {
					continue gameLoop
				}
			}
		}

		validGameIdxSum += (gameIdx + 1)
	}

	fmt.Println(validGameIdxSum)
}
