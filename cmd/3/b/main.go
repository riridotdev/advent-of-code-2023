package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	rows := strings.Split(inputString, "\n")

	visited := make([][]bool, len(rows))
	for i := range visited {
		visited[i] = make([]bool, len(rows[0]))
	}

	gearRatioSum := 0
	for y := range rows {
		for x := range rows[y] {
			currentChar := rune(rows[y][x])
			if !unicode.IsDigit(currentChar) && currentChar != '.' {
				directions := [][2]int{
					{1, 0},
					{1, -1},
					{0, -1},
					{-1, -1},
					{-1, 0},
					{-1, 1},
					{0, 1},
					{1, 1},
				}

				numbers := []int{}
				for _, direction := range directions {
					newX, newY := x+direction[0], y+direction[1]
					newChar := rune(rows[newY][newX])

					if visited[newY][newX] {
						continue
					}
					if !unicode.IsDigit(newChar) {
						continue
					}

					visited[newY][newX] = true

					numberString := extractNumberString(rows, newX, newY, visited)
					number, _ := strconv.Atoi(numberString)

					numbers = append(numbers, number)
				}

				if len(numbers) == 2 {
					gearRatioSum += numbers[0] * numbers[1]
				}
			}
		}
	}

	fmt.Println(gearRatioSum)
}

func extractNumberString(rows []string, x, y int, visited [][]bool) string {
	start, end := x, x+1

	queue := [][2]int{{x, y}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		currentX := current[0]

		xDeltas := [2]int{1, -1}
		for _, xDelta := range xDeltas {
			newX := currentX + xDelta
			if newX < len(rows[0]) && newX >= 0 && !visited[y][newX] && unicode.IsDigit(rune(rows[y][newX])) {
				queue = append(queue, [2]int{newX, y})
				start = min(start, newX)
				end = max(end, newX+1)
				visited[y][newX] = true
			}
		}
	}

	return rows[y][start:end]
}
