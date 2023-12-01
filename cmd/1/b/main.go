package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stringToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	lines := strings.Split(inputString, "\n")

	total := 0
	for _, line := range lines {
        seen := map[[2]int]struct{}{}
        nums := parseNumbers(line, 0, len(line), seen)
		total += (nums[0] * 10) + nums[len(nums)-1]
	}

	fmt.Println(total)
}

func parseNumbers(s string, i, j int, seen map[[2]int]struct{}) []int {
	if _, ok := seen[[2]int{i, j}]; ok {
		return []int{}
	}
	if j-i == 0 {
		return []int{}
	}

	seen[[2]int{i, j}] = struct{}{}

	if j-i == 1 {
        val, err := strconv.Atoi(s[i:j])
		if err != nil {
			return []int{}
		}
		return []int{val}
	}

    if val, ok := stringToDigit[s[i:j]]; ok {
		return []int{val}
	}

	results := []int{}
	results = append(results, parseNumbers(s, i, j-1, seen)...)
	results = append(results, parseNumbers(s, i+1, j, seen)...)

	return results
}
