package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	possibleArrangements := 0

	records := strings.Split(inputString, "\n")
	for _, record := range records {
		recordParts := strings.Fields(record)

		springMap := recordParts[0]
		groupingCounts := []int{}

		for _, groupingCountStr := range strings.Split(recordParts[1], ",") {
			groupingCount, _ := strconv.Atoi(groupingCountStr)
			groupingCounts = append(groupingCounts, groupingCount)
		}

		results := []string{}
		findPermutations(springMap, 0, groupingCounts, 0, &results)
		possibleArrangements += len(results)
	}

	print(possibleArrangements)
}

func findPermutations(springMap string, pos int, groupings []int, groupingIdx int, results *[]string) {
	if pos == len(springMap) {
		for _, grouping := range groupings {
			if grouping != 0 {
				return
			}
		}
		*results = append(*results, springMap)
		return
	}

	if springMap[pos] == '#' {
		if groupingIdx == len(groupings) {
			return
		}
		if groupings[groupingIdx] == 0 {
			return
		}
		groupings[groupingIdx] -= 1
		findPermutations(springMap, pos+1, groupings, groupingIdx, results)
		groupings[groupingIdx] += 1
		return
	}
	if springMap[pos] == '.' {
		if pos > 0 && springMap[pos-1] == '#' {
			if groupings[groupingIdx] != 0 {
				return
			}
			findPermutations(springMap, pos+1, groupings, groupingIdx+1, results)
			return
		}
		findPermutations(springMap, pos+1, groupings, groupingIdx, results)
		return
	}
	if springMap[pos] == '?' {
		springMapWithDamaged := []byte(springMap)
		springMapWithDamaged[pos] = '#'
		findPermutations(string(springMapWithDamaged), pos, groupings, groupingIdx, results)

		springMapWithOperational := []byte(springMap)
		springMapWithOperational[pos] = '.'
		findPermutations(string(springMapWithOperational), pos, groupings, groupingIdx, results)
		return
	}

	panic("invalid spring map")
}
