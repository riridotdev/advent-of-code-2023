package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Left
	Down
	Right
	None
)

var deltas = [4][2]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

var validConnections = map[byte][]Direction{
	'|': {Up, Down},
	'-': {Left, Right},
	'F': {Down, Right},
	'7': {Down, Left},
	'L': {Up, Right},
	'J': {Up, Left},
}

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	rowStrings := strings.Split(inputString, "\n")

	grid := [][]byte{}
	for _, row := range rowStrings {
		grid = append(grid, []byte(row))
	}

	startX, startY := findStart(grid)
	replaceStart(grid, startX, startY)

	visited := map[[2]int]struct{}{
		{startX, startY}: {},
	}

	loopGrid := make([][]byte, len(grid))
	for i := range grid {
		loopGrid[i] = make([]byte, len(grid[0]))
		for j := range loopGrid[i] {
			loopGrid[i][j] = '.'
		}
	}

	queue := [][2]int{{startX, startY}}
	for len(queue) != 0 {
		for i := 0; i < len(queue); i++ {
			current := queue[0]
			queue = queue[1:]

			for _, delta := range deltas {
				currentX, currentY := current[0], current[1]
				loopGrid[currentY][currentX] = grid[currentY][currentX]
				targetX, targetY := currentX+delta[0], currentY+delta[1]
				if targetX >= len(grid[0]) || targetX < 0 || targetY >= len(grid) || targetY < 0 {
					continue
				}
				if _, ok := visited[[2]int{targetX, targetY}]; ok {
					continue
				}
				if connects(grid, currentX, currentY, targetX, targetY) {
					visited[[2]int{targetX, targetY}] = struct{}{}
					queue = append(queue, [2]int{targetX, targetY})
				}
			}
		}
	}

	count := 0
	for y := range loopGrid {
		crosses := 0
		openDir := None
		for x := range loopGrid[y] {
			switch loopGrid[y][x] {
			case '.':
				if crosses != 0 && crosses%2 != 0 {
					count += 1
				}
			case '|':
				crosses += 1
			case 'F':
				switch openDir {
				case None:
					openDir = Down
				case Up:
					openDir = None
					crosses += 1
				case Down:
					openDir = None
				}
			case '7':
				switch openDir {
				case None:
					openDir = Down
				case Up:
					openDir = None
					crosses += 1
				case Down:
					openDir = None
				}
			case 'L':
				switch openDir {
				case None:
					openDir = Up
				case Up:
					openDir = None
				case Down:
					openDir = None
					crosses += 1
				}
			case 'J':
				switch openDir {
				case None:
					openDir = Up
				case Up:
					openDir = None
				case Down:
					openDir = None
					crosses += 1
				}
			}
		}
	}

	fmt.Println(count)
}

func findStart(grid [][]byte) (int, int) {
	for y, row := range grid {
		for x, tile := range row {
			if tile != 'S' {
				continue
			}
			return x, y
		}
	}
	panic("Start not found")
}

func replaceStart(grid [][]byte, startX, startY int) {
	connections := []Direction{}
	for _, delta := range deltas {
		targetX, targetY := startX+delta[0], startY+delta[1]
		if targetX >= len(grid[0]) || targetX < 0 || targetY >= len(grid) || targetY < 0 {
			continue
		}
		direction := getDirection(targetX, targetY, startX, startY)
		for _, validConnection := range validConnections[grid[targetY][targetX]] {
			if validConnection == direction {
				connections = append(connections, reverseDirection(validConnection))
			}
		}
	}

	replacementTile := tileByConnectionDirections(connections)

	grid[startY][startX] = replacementTile
}

func connects(grid [][]byte, currentX, currentY, targetX, targetY int) bool {
	source := grid[currentY][currentX]
	target := grid[targetY][targetX]

	direction := getDirection(currentX, currentY, targetX, targetY)
	var sourceValid bool
	for _, connection := range validConnections[source] {
		if connection == direction {
			sourceValid = true
			break
		}
	}

	oppositeDir := reverseDirection(direction)
	var targetValid bool
	for _, connection := range validConnections[target] {
		if connection == oppositeDir {
			targetValid = true
			break
		}
	}

	return sourceValid && targetValid
}

func getDirection(currentX, currentY, targetX, targetY int) Direction {
	if targetX > currentX {
		return Right
	}
	if targetX < currentX {
		return Left
	}
	if targetY > currentY {
		return Down
	}
	if targetY < currentY {
		return Up
	}
	panic("Invalid movement")
}

func tileByConnectionDirections(dirs []Direction) byte {
outer:
	for tile, connections := range validConnections {
		for _, dir := range dirs {
			if !in(dir, connections) {
				continue outer
			}
		}
		return tile
	}
	panic("Invalid connection pairing")
}

func reverseDirection(dir Direction) Direction {
	return Direction((int(dir) + 2) % 4)
}

func in[T comparable](val T, arr []T) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
