package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	games := strings.Split(inputString, "\n")

	powerSum := 0
	for _, game := range games {
		game = strings.SplitAfter(game, ":")[1]
		combinations := strings.Split(game, ";")

		cubesRequired := map[string]int{}

		for _, combination := range combinations {
			colourCounts := strings.Split(combination, ",")

			for _, colourCount := range colourCounts {
				parts := strings.Split(strings.TrimSpace(colourCount), " ")

				countString := parts[0]
				count, _ := strconv.Atoi(countString)

				colour := parts[1]

				cubesRequired[colour] = max(cubesRequired[colour], count)
			}
		}

		power := 1
		for _, count := range cubesRequired {
			power *= count
		}

		powerSum += power
	}

	fmt.Println(powerSum)
}
