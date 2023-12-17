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
	matchingCount := 0
	for matchingCount <= row {
		leftRow := row - matchingCount
		rightRow := row + matchingCount + 1
		if leftRow < 0 || rightRow == len(grid) {
			return true
		}
		if !rowsMatch(grid, leftRow, rightRow) {
			return false
		}
		matchingCount += 1
	}
	return true
}

func rowsMatch(grid [][]byte, a, b int) bool {
	for i := range grid[0] {
		if grid[a][i] != grid[b][i] {
			return false
		}
	}
	return true
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
	matchingCount := 0
	for matchingCount <= col {
		leftCol := col - matchingCount
		rightCol := col + matchingCount + 1
		if leftCol < 0 || rightCol == len(grid[0]) {
			return true
		}
		if !colsMatch(grid, leftCol, rightCol) {
			return false
		}
		matchingCount += 1
	}
	return true
}

func colsMatch(grid [][]byte, a, b int) bool {
	for i := range grid {
		if grid[i][a] != grid[i][b] {
			return false
		}
	}
	return true
}
