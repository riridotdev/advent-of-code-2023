package main

import (
	"fmt"
	"os"
)

type GalaxyPair [2]struct {
	x int
	y int
}

type GalaxyDistances map[GalaxyPair]int

func (gd GalaxyDistances) Contains(gp GalaxyPair) bool {
	if _, ok := gd[gp]; ok {
		return true
	}
	if _, ok := gd[GalaxyPair{gp[1], gp[0]}]; ok {
		return true
	}
	return false
}

func main() {
	inputBytes, _ := os.ReadFile("input")

	grid := [][]byte{}

	current := 0
	for len(inputBytes) != 0 {
		if current == len(inputBytes) || inputBytes[current] == '\n' {
			grid = append(grid, inputBytes[:current])
			inputBytes = inputBytes[current+1:]
			current = 0
		}
		current += 1
	}

	expandRows := []int{}
	expandCols := []int{}

	for i := range grid {
		if rowEmpty(grid, i) {
			expandRows = append(expandRows, i)
		}
	}
	for i := range grid[0] {
		if colEmpty(grid, i) {
			expandCols = append(expandCols, i)
		}
	}

	tempGrid := [][]byte{}
	colOffset := 0
	for rowIdx := range grid {
		row := make([]byte, len(grid[rowIdx])+len(expandCols))
		for colIdx := range grid[rowIdx] {
			current := grid[rowIdx][colIdx]
			row[colIdx+colOffset] = current
			for _, expandCol := range expandCols {
				if colIdx == expandCol {
					colOffset += 1
					row[colIdx+colOffset] = current
					break
				}
			}
		}
		colOffset = 0

		tempGrid = append(tempGrid, row)
		for _, expandRow := range expandRows {
			if rowIdx == expandRow {
				tempGrid = append(tempGrid, row)
				break
			}
		}
	}
	grid = tempGrid

	pairDistances := GalaxyDistances{}
	distanceCount := 0
	for y := range grid {
		for x := range grid[y] {
			current := grid[y][x]
			if current == '.' {
				continue
			}
			distanceCount += totalDistance(grid, x, y, pairDistances)
		}
	}

	fmt.Println(distanceCount)
}

func totalDistance(grid [][]byte, startX, startY int, results GalaxyDistances) int {
	totalDistance := 0
	steps := 1
	visited := map[[2]int]struct{}{
		{startX, startY}: {},
	}
	queue := [][2]int{{startX, startY}}
	for len(queue) != 0 {
		for i := len(queue); i != 0; i-- {
			current := queue[0]
			queue = queue[1:]

			deltas := [4][2]int{
				{1, 0},
				{0, -1},
				{-1, 0},
				{0, 1},
			}
			for _, delta := range deltas {
				newX, newY := current[0]+delta[0], current[1]+delta[1]
				if newX >= len(grid[0]) || newX < 0 ||
					newY >= len(grid) || newY < 0 {
					continue
				}

				if _, ok := visited[[2]int{newX, newY}]; ok {
					continue
				}
				visited[[2]int{newX, newY}] = struct{}{}

				next := grid[newY][newX]
				if next != '.' {
					resultsIdx := GalaxyPair{{startX, startY}, {newX, newY}}
					if results.Contains(resultsIdx) {
						continue
					}
					results[resultsIdx] = steps
					totalDistance += steps
				}

				queue = append(queue, [2]int{newX, newY})
			}
		}
		steps += 1
	}

	return totalDistance
}

func rowEmpty(grid [][]byte, i int) bool {
	for j := 0; j < len(grid[i]); j++ {
		if grid[i][j] != '.' {
			return false
		}
	}
	return true
}

func colEmpty(grid [][]byte, i int) bool {
	for j := 0; j < len(grid); j++ {
		if grid[j][i] != '.' {
			return false
		}
	}
	return true
}

func printGrid(grid [][]byte) {
	galaxyCount := 0
	for y := range grid {
		for x := range grid[y] {
			current := grid[y][x]
			if current == '#' {
				galaxyCount += 1
				fmt.Printf("%d", galaxyCount)
			} else {
				fmt.Printf("%c", current)
			}
		}
		print("\n")
	}
}
