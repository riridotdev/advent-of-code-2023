package main

import (
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	total := 0
	for _, step := range strings.Split(inputString, ",") {
		hash := 0
		for _, char := range step {
			hash += int(char)
			hash *= 17
			hash %= 256
		}
		total += hash
	}

	println(total)
}
