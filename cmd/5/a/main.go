package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	sections := strings.Split(inputString, "\n\n")

	seedsString := strings.Split(sections[0], ":")[1]
	seedStrings := strings.Fields(seedsString)

	translationMaps := []TranslationMap{}
	translationMapStrings := sections[1:]
	for _, translationMapString := range translationMapStrings {
		translationMaps = append(translationMaps, NewTranslationMap(translationMapString))
	}

	lowest := math.MaxInt
	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		for _, translationMap := range translationMaps {
			seed = translationMap.translate(seed)
		}
		lowest = min(lowest, seed)
	}

	fmt.Println(lowest)
}

type TranslationMap []TranslationRange

func NewTranslationMap(source string) TranslationMap {
	translationMap := TranslationMap{}
	lines := strings.Split(source, "\n")
	for _, line := range lines[1:] {
		translationMap = append(translationMap, NewTranslationRange(line))
	}
	return translationMap
}

func (tm TranslationMap) translate(i int) int {
	for _, tr := range tm {
		if (i >= tr.Source) && (i <= tr.Source+tr.Range) {
			return (i - tr.Source) + tr.Destination
		}
	}
	return i
}

type TranslationRange struct {
	Source      int
	Destination int
	Range       int
}

func NewTranslationRange(source string) TranslationRange {
	fieldStrings := strings.Fields(source)

	fields := [3]int{}
	for i, fieldString := range fieldStrings {
		field, _ := strconv.Atoi(fieldString)
		fields[i] = field
	}

	return TranslationRange{
		Destination: fields[0],
		Source:      fields[1],
		Range:       fields[2],
	}
}
