package main

import (
	"bufio"
	"os"
)

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

	tilt(grid)

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

func tilt(grid [][]byte) {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != 'O' {
				continue
			}
			rockRow := row
			for rockRow >= 0 && rockRow != 0 && grid[rockRow-1][col] == '.' {
				grid[rockRow][col] = '.'
				rockRow -= 1
				grid[rockRow][col] = 'O'
			}
		}
	}
}

func drawGrid(grid [][]byte) {
	for _, row := range grid {
		println(string(row))
	}
}
