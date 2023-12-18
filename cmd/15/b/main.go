package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	boxes := [256][]lens{}
	for i := range boxes {
		boxes[i] = []lens{}
	}

	for _, step := range strings.Split(inputString, ",") {
		parts := strings.FieldsFunc(step, func(r rune) bool {
			return r == '=' || r == '-'
		})

		label := parts[0]

		if len(parts) == 2 {
			focalLength, _ := strconv.Atoi(parts[1])
			handleEquals(&boxes, label, focalLength)
		}
		if len(parts) == 1 {
			handleMinus(&boxes, label)
		}
	}

	focusingPower := 0
	for i, box := range boxes {
		for j, lens := range box {
			focusingPower += (1 + i) * (j + 1) * lens.focalLength
		}
	}

	println(focusingPower)
}

type lens struct {
	label       string
	focalLength int
}

func handleEquals(boxes *[256][]lens, label string, focalLength int) {
	boxIdx := hashString(label)
	for i, lens := range boxes[boxIdx] {
		if lens.label == label {
			boxes[boxIdx][i].focalLength = focalLength
			return
		}
	}
	boxes[boxIdx] = append(boxes[boxIdx], lens{label, focalLength})
}

func handleMinus(boxes *[256][]lens, label string) {
	boxIdx := hashString(label)
	replaceIdx := -1
	for i, lens := range boxes[boxIdx] {
		if lens.label == label {
			replaceIdx = i + 1
		}
	}
	if replaceIdx == -1 {
		return
	}
	for replaceIdx != len(boxes[boxIdx]) {
		boxes[boxIdx][replaceIdx-1] = boxes[boxIdx][replaceIdx]
		replaceIdx += 1
	}
	boxes[boxIdx] = boxes[boxIdx][:len(boxes[boxIdx])-1]
}

func hashString(s string) int {
	hash := 0
	for _, char := range s {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}
