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

	sequenceStrings := strings.Split(inputString, "\n")

	sequences := [][]int{}
	for _, sequenceString := range sequenceStrings {
		numbers := strings.Fields(sequenceString)
		sequence := []int{}
		for _, number := range numbers {
			val, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			sequence = append(sequence, val)
		}
		sequences = append(sequences, sequence)
	}

	total := 0
	for _, sequence := range sequences {
		total += nextVal(sequence)
	}

	fmt.Println(total)
}

func nextVal(sequence []int) int {
	if allZero(sequence) {
		return 0
	}

	diffSeq := []int{}
	for i := 0; i < len(sequence)-1; i++ {
		diffSeq = append(diffSeq, sequence[i+1]-sequence[i])
	}

	return sequence[len(sequence)-1] + nextVal(diffSeq)
}

func allZero(sequence []int) bool {
	for _, val := range sequence {
		if val != 0 {
			return false
		}
	}
	return true
}
