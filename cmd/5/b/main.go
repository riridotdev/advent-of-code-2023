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

	seeds := []int{}
	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	lowest := math.MaxInt
	for len(seeds) != 0 {
		start := seeds[0]
		length := seeds[1]
		seedRanges := []SeedRange{{Start: start, End: (start + length) - 1}}
		seeds = seeds[2:]

		for _, translationMap := range translationMaps {
			seedRanges = translationMap.translate(seedRanges)
		}

		for _, seedRange := range seedRanges {
			lowest = min(lowest, seedRange.Start)
		}
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

func (tm TranslationMap) translate(seedRanges []SeedRange) []SeedRange {
	translatedRanges := []SeedRange{}
	for _, tr := range tm {
		next := []SeedRange{}
		for _, seedRange := range seedRanges {
			before, translated, after := tr.translate(seedRange)
			if before != nil {
				next = append(next, *before)
			}
			if after != nil {
				next = append(next, *after)
			}
			if translated != nil {
				translatedRanges = append(translatedRanges, *translated)
			}
		}
		seedRanges = next
	}
	translatedRanges = append(translatedRanges, seedRanges...)
	return translatedRanges
}

type TranslationRange struct {
	SourceRange      SeedRange
	DestinationRange SeedRange
}

func NewTranslationRange(rangeString string) TranslationRange {
	fieldStrings := strings.Fields(rangeString)

	fields := [3]int{}
	for i, fieldString := range fieldStrings {
		field, _ := strconv.Atoi(fieldString)
		fields[i] = field
	}
	destination, source, length := fields[0], fields[1], fields[2]

	return TranslationRange{
		SourceRange:      SeedRange{Start: source, End: (source + length) - 1},
		DestinationRange: SeedRange{Start: destination, End: (destination + length) - 1},
	}
}

func (tr TranslationRange) translate(seedRange SeedRange) (before, translated, after *SeedRange) {
	if seedRange.Start < tr.SourceRange.Start {
		before = &SeedRange{Start: seedRange.Start, End: min(seedRange.End, tr.SourceRange.Start-1)}
	}
	if seedRange.End > tr.SourceRange.End {
		after = &SeedRange{Start: max(seedRange.Start, tr.SourceRange.End+1), End: seedRange.End}
	}

	innerRangeStart := max(seedRange.Start, tr.SourceRange.Start)
	innerRangeEnd := min(seedRange.End, tr.SourceRange.End)
	offset := innerRangeStart - tr.SourceRange.Start
	length := (innerRangeEnd - innerRangeStart) + 1

	if length <= 0 {
		return
	}

	translated = &SeedRange{
		Start: tr.DestinationRange.Start + offset,
		End:   tr.DestinationRange.Start + offset + length - 1,
	}

	return
}

type SeedRange struct {
	Start int
	End   int
}
