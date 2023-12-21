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

	start := findStart(grid)
	grid[start.y][start.x] = '.'

	const steps = 64

	queue := []point{start}
	for i := 0; i < steps; i++ {
		currentLocations := map[point]struct{}{}
		for i := len(queue); i > 0; i-- {
			current := queue[0]
			queue = queue[1:]

			directions := []point{
				{1, 0},
				{0, 1},
				{-1, 0},
				{0, -1},
			}
			for _, dir := range directions {
				nextX, nextY := current.x+dir.x, current.y+dir.y
				if nextX >= len(grid[0]) || nextX < 0 || nextY >= len(grid) || nextY < 0 {
					continue
				}
				if grid[nextY][nextX] == '#' {
					continue
				}
				if _, ok := currentLocations[point{nextX, nextY}]; ok {
					continue
				}
				currentLocations[point{nextX, nextY}] = struct{}{}
				queue = append(queue, point{nextX, nextY})
			}
		}
	}

	println(len(queue))
}

func findStart(grid [][]byte) point {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != 'S' {
				continue
			}
			return point{x, y}
		}
	}
	panic("failed to find start in grid")
}

type point struct {
	x, y int
}
