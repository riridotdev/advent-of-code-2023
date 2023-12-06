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

	inputParts := strings.Split(inputString, "\n")

	timesString := strings.TrimSpace(strings.Split(inputParts[0], ":")[1])
	distancesString := strings.TrimSpace(strings.Split(inputParts[1], ":")[1])

	time := 0
	for _, timeString := range strings.Fields(timesString) {
		timeInt, _ := strconv.Atoi(timeString)
		for i := 0; i < len(timeString); i++ {
			time *= 10
		}
		time += timeInt
	}

	distance := 0
	for _, distanceString := range strings.Fields(distancesString) {
		distanceInt, _ := strconv.Atoi(distanceString)
		for i := 0; i < len(distanceString); i++ {
			distance *= 10
		}
		distance += distanceInt
	}

	l := 0
	r := time
	var leftBisect int
	for l < r {
		holdTime := (l + r) / 2
		travelled := holdTime * (time - holdTime)

		if travelled > distance {
			leftBisect = holdTime
			r = holdTime
			continue
		}

		l = holdTime + 1
	}

	l = 0
	r = time
	var rightBisect int
	for l < r {
		holdTime := (l + r) / 2
		travelled := holdTime * (time - holdTime)

		if travelled > distance {
			rightBisect = holdTime
			l = holdTime + 1
			continue
		}

		r = holdTime
	}

	possibleVictories := (rightBisect - leftBisect) + 1

	fmt.Println(possibleVictories)
}
