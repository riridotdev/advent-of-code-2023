package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	patterns := strings.Split(inputString, "\n\n")

	verticalCount := 0
	horizontalCount := 0

	for _, pattern := range patterns {
		grid := [][]byte{}
		for _, row := range strings.Split(pattern, "\n") {
			grid = append(grid, []byte(row))
		}

		verticalCount += findVerticalReflection(grid)
		horizontalCount += findHorizontalReflection(grid)
	}

	fmt.Println((horizontalCount * 100) + verticalCount)
}

func findHorizontalReflection(grid [][]byte) int {
	for i := 0; i < len(grid)-1; i++ {
		if reflectsAroundRow(grid, i) {
			return i + 1
		}
	}
	return 0
}

func reflectsAroundRow(grid [][]byte, row int) bool {
	remainingFixes := 1
	matchingCount := 0
	for matchingCount <= row {
		leftRow := row - matchingCount
		rightRow := row + matchingCount + 1
		if leftRow < 0 || rightRow == len(grid) {
			break
		}
		diff := rowsDiff(grid, leftRow, rightRow)
		if diff > remainingFixes {
			return false
		}
		remainingFixes -= diff
		matchingCount += 1
	}
	return remainingFixes == 0
}

func rowsDiff(grid [][]byte, a, b int) int {
	count := 0
	for i := range grid[0] {
		if grid[a][i] != grid[b][i] {
			count += 1
		}
	}
	return count
}

func findVerticalReflection(grid [][]byte) int {
	for i := 0; i < len(grid[0])-1; i++ {
		if reflectsAroundCol(grid, i) {
			return i + 1
		}
	}
	return 0
}

func reflectsAroundCol(grid [][]byte, col int) bool {
	remainingFixes := 1
	matchingCount := 0
	for matchingCount <= col {
		leftCol := col - matchingCount
		rightCol := col + matchingCount + 1
		if leftCol < 0 || rightCol == len(grid[0]) {
			break
		}
		diff := colsDiff(grid, leftCol, rightCol)
		if diff > remainingFixes {
			return false
		}
		remainingFixes -= diff
		matchingCount += 1
	}
	return remainingFixes == 0
}

func colsDiff(grid [][]byte, a, b int) int {
	count := 0
	for i := range grid {
		if grid[i][a] != grid[i][b] {
			count += 1
		}
	}
	return count
}
