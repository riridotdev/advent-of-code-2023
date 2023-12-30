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

	start := point{1, 0}
	end := point{len(grid) - 2, len(grid) - 1}

	println(longestPath(grid, start, end, map[point]struct{}{}))
}

func longestPath(grid [][]byte, pos, end point, seen map[point]struct{}) int {
	if pos == end {
		return 0
	}

	if pos.x >= len(grid[0]) || pos.x < 0 || pos.y >= len(grid) || pos.y < 0 {
		return -1
	}
	if grid[pos.y][pos.x] == '#' {
		return -1
	}

	if _, ok := seen[pos]; ok {
		return -1
	}
	seen[pos] = struct{}{}

	directions := []point{
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 0},
	}

	maxChild := -1
	for _, dir := range directions {
		pathLen := longestPath(grid, point{pos.x + dir.x, pos.y + dir.y}, end, seen)
		maxChild = max(pathLen, maxChild)
	}

	delete(seen, pos)

	if maxChild == -1 {
		return -1
	}
	return 1 + maxChild
}

type point struct {
	x, y int
}
