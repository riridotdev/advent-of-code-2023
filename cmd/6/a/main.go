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

	times := []int{}
	for _, timeString := range strings.Fields(timesString) {
		time, _ := strconv.Atoi(timeString)
		times = append(times, time)
	}

	distances := []int{}
	for _, distanceString := range strings.Fields(distancesString) {
		distance, _ := strconv.Atoi(distanceString)
		distances = append(distances, distance)
	}

	possibleVictoriesProduct := 1
	for raceIdx := range times {
		time := times[raceIdx]
		distance := distances[raceIdx]

		possibleVictories := 0
		for holdTime := 1; holdTime < time; holdTime++ {
			travelled := holdTime * (time - holdTime)
			if travelled <= distance {
				continue
			}
			possibleVictories += 1
		}

		possibleVictoriesProduct *= possibleVictories
	}

	fmt.Println(possibleVictoriesProduct)
}
