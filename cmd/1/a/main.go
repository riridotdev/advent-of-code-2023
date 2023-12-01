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

	lines := strings.Split(inputString, "\n")

	total := 0
	for _, line := range lines {
		firstNum := -1
		lastNum := 0
		for _, char := range line {
			val, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			if firstNum == -1 {
				firstNum = val
			}
			lastNum = val
		}
		total += (firstNum * 10) + lastNum
	}

	fmt.Println(total)
}
