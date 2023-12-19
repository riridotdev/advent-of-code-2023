package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

var directions = []point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	xMin, xMax, yMin, yMax := 0, 0, 0, 0

	position := point{0, 0}

	digMap := map[point]struct{}{
		position: {},
	}
	for _, command := range strings.Split(inputString, "\n") {
		var direction byte
		var countByte string
		fmt.Sscanf(command, "%c %s", &direction, &countByte)

		vector := vectorForDirection(direction)

		count, _ := strconv.Atoi(string(countByte))
		for count != 0 {
			position.x += vector.x
			position.y += vector.y
			digMap[position] = struct{}{}
			count -= 1
		}

		xMin = min(position.x, xMin)
		xMax = max(position.x, xMax)
		yMin = min(position.y, yMin)
		yMax = max(position.y, yMax)
	}

	gridX := (xMax - xMin) + 1
	gridY := (yMax - yMin) + 1

	xOffset := xMin
	yOffset := yMin

	grid := make([][]byte, gridY)
	for i := range grid {
		grid[i] = make([]byte, gridX)
	}
	for y := range grid {
		for x := range grid[y] {
			if _, ok := digMap[point{x + xOffset, y + yOffset}]; ok {
				grid[y][x] = '#'
			} else {
				grid[y][x] = '.'
			}
		}
	}

	visited := map[point]struct{}{}
	for y := range grid {
		inside := false
		for x := range grid[y] {
			if grid[y][x] == '#' {
				if x == 0 {
					inside = true
					continue
				}
				if grid[y][x-1] == '.' {
					inside = !inside
				}
			}
			if inside {
				if isContained(grid, x, y, visited, map[point]struct{}{}) {
					markContained(grid, x, y)
				}
			}
		}
	}

	count := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '#' {
				count += 1
			}
		}
	}

	println(count)
}

func isContained(grid [][]byte, x, y int, visited map[point]struct{}, visiting map[point]struct{}) bool {
	if x >= len(grid[0]) || x < 0 || y >= len(grid) || y < 0 {
		return false
	}

	if _, ok := visited[point{x, y}]; ok {
		return false
	}

	if grid[y][x] == '#' {
		return true
	}

	for _, direction := range directions {
		nextX, nextY := x+direction.x, y+direction.y
		if _, ok := visiting[point{nextX, nextY}]; ok {
			continue
		}
		visiting[point{nextX, nextY}] = struct{}{}
		if !isContained(grid, x+direction.x, y+direction.y, visited, visiting) {
			visited[point{x, y}] = struct{}{}
			return false
		}
	}

	visited[point{x, y}] = struct{}{}
	return true
}

func markContained(grid [][]byte, x, y int) {
	if x >= len(grid[0]) || x < 0 || y >= len(grid) || y < 0 {
		return
	}

	if grid[y][x] == '#' {
		return
	}

	grid[y][x] = '#'

	for _, direction := range directions {
		markContained(grid, x+direction.x, y+direction.y)
	}

}

func vectorForDirection(dir byte) point {
	switch dir {
	case 'U':
		return point{0, -1}
	case 'D':
		return point{0, 1}
	case 'L':
		return point{-1, 0}
	case 'R':
		return point{1, 0}
	}
	panic("invalid direction")
}
