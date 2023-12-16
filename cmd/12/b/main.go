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

	possibleArrangements := 0

	records := strings.Split(inputString, "\n")
	for _, record := range records {
		recordParts := strings.Fields(record)

		springMap := recordParts[0]
		groupingsString := recordParts[1]

		repeats := []string{}
		for i := 0; i < 5; i++ {
			repeats = append(repeats, springMap)
		}
		springMap = strings.Join(repeats, "?")

		repeats = []string{}
		for i := 0; i < 5; i++ {
			repeats = append(repeats, groupingsString)
		}
		groupingsString = strings.Join(repeats, ",")

		groupings := []int{}
		for _, groupingCountStr := range strings.Split(groupingsString, ",") {
			groupingCount, _ := strconv.Atoi(groupingCountStr)
			groupings = append(groupings, groupingCount)
		}

		possibleArrangements += findPermutations(springMap, groupings, '.', map[string]int{})
	}

	print(possibleArrangements)
}

func memoIndex(springMap string, groupings []int, previous byte) string {
	return fmt.Sprintf("%s%v%c", springMap, groupings, previous)
}

func findPermutations(springMap string, groupings []int, previous byte, memo map[string]int) int {
	if springMap == "" {
		for _, grouping := range groupings {
			if grouping != 0 {
				return 0
			}
		}
		return 1
	}

	memoIndex := memoIndex(springMap, groupings, previous)
	if val, ok := memo[memoIndex]; ok {
		return val
	}

	var result int

	if springMap[0] == '#' {
		if len(groupings) == 0 || groupings[0] == 0 {
			return 0
		}
		groupings = decrementGroupings(groupings)
		result = findPermutations(springMap[1:], groupings, springMap[0], memo)
	}
	if springMap[0] == '.' {
		if previous == '#' {
			if groupings[0] != 0 {
				return 0
			}
			result = findPermutations(springMap[1:], groupings[1:], springMap[0], memo)
		} else {
			result = findPermutations(springMap[1:], groupings, springMap[0], memo)
		}
	}
	if springMap[0] == '?' {
		mapBytes := []byte(springMap)

		mapBytes[0] = '#'
		result += findPermutations(string(mapBytes), groupings, previous, memo)

		mapBytes[0] = '.'
		result += findPermutations(string(mapBytes), groupings, previous, memo)
	}

	memo[memoIndex] = result

	return result
}

func prependResults(results [][]byte, b byte) [][]byte {
	for i, result := range results {
		results[i] = append([]byte{b}, result...)
	}
	return results
}

func decrementGroupings(groupings []int) []int {
	result := make([]int, len(groupings))
	copy(result, groupings)
	result[0] -= 1
	return result
}
