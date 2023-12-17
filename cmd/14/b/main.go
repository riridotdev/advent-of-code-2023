package main

import (
	"bufio"
	"os"
	"strings"
)

type direction [2]int

var (
	north direction = [2]int{-1, 0}
	west  direction = [2]int{0, -1}
	south direction = [2]int{1, 0}
	east  direction = [2]int{0, 1}
)

var directions = []direction{
	north,
	west,
	south,
	east,
}

func main() {
	inputFile, _ := os.Open("input")
	input := bufio.NewReader(inputFile)

	grid := [][]byte{}
	for {
		line, err := input.ReadBytes('\n')
		if err != nil {
			break
		}
		grid = append(grid, line[:len(line)-1])
	}

	spinCount := 1000000000

	firstSeen := map[string]int{}
	cycleLength := map[string]int{}

	for i := 0; i <= spinCount; i++ {
		gridString := gridToString(grid)

		length, ok := cycleLength[gridString]
		if !ok {
			if seenAt, ok := firstSeen[gridString]; ok {
				length = i - seenAt
				cycleLength[gridString] = length
			} else {
				firstSeen[gridString] = i
			}
		}

		if length > 0 && (i+length) < spinCount {
			for (i + length) < spinCount {
				i += length
			}
			continue
		}

		for _, direction := range directions {
			tilt(grid, direction)
		}
	}

	totalLoad := 0
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == 'O' {
				totalLoad += len(grid) - row
			}
		}
	}

	println(totalLoad)
}

func tilt(grid [][]byte, dir direction) {
	rowDelta := dir[0]
	colDelta := dir[1]

	rowStart := max(0, rowDelta*len(grid)-1)
	colStart := max(0, colDelta*len(grid[0])-1)

	rowIterator := 1
	if rowStart > 0 {
		rowIterator = -1
	}
	colIterator := 1
	if colStart > 0 {
		colIterator = -1
	}

	for row := rowStart; row >= 0 && row < len(grid); row += rowIterator {
		for col := colStart; col >= 0 && col < len(grid[0]); col += colIterator {
			if grid[row][col] != 'O' {
				continue
			}
			rockRow := row
			rockCol := col
			for {
				nextRow := rockRow + rowDelta
				nextCol := rockCol + colDelta
				if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
					break
				}
				if grid[nextRow][nextCol] != '.' {
					break
				}
				grid[rockRow][rockCol] = '.'
				rockRow = nextRow
				rockCol = nextCol
				grid[rockRow][rockCol] = 'O'
			}

		}
	}
}

func gridToString(grid [][]byte) string {
	builder := strings.Builder{}
	for _, row := range grid {
		builder.Write(row)
	}
	return builder.String()
}

func drawGrid(grid [][]byte) {
	for _, row := range grid {
		println(string(row))
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
