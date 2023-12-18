package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction [2]int

var (
	North = direction{0, -1}
	South = direction{0, 1}
	East  = direction{1, 0}
	West  = direction{-1, 0}
)

type beam struct {
	x, y int
	dir  direction
}

func (b beam) nextPos() (int, int) {
	return b.x + b.dir[0], b.y + b.dir[1]
}

func main() {
	inputFile, _ := os.Open("input")
	input := bufio.NewReader(inputFile)

	grid := [][]byte{}
	for {
		row, err := input.ReadBytes('\n')
		if err != nil {
			break
		}
		grid = append(grid, row[:len(row)-1])
	}

	illuminated := map[[2]int]struct{}{}

	seen := map[beam]struct{}{}

	queue := []beam{{-1, 0, East}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if _, ok := seen[current]; ok {
			continue
		}
		seen[current] = struct{}{}

		queue = append(queue, stepBeam(grid, current)...)

		if current.x >= len(grid[0]) || current.x < 0 || current.y >= len(grid) || current.y < 0 {
			continue
		}

		illuminated[[2]int{current.x, current.y}] = struct{}{}
	}

	println(len(illuminated))
}

func stepBeam(grid [][]byte, current beam) []beam {
	nextX, nextY := current.nextPos()
	if nextX >= len(grid[0]) || nextX < 0 || nextY >= len(grid) || nextY < 0 {
		return nil
	}

	directions := []direction{}
	switch grid[nextY][nextX] {
	case '.':
		directions = []direction{current.dir}
	case '|':
		if current.dir == North || current.dir == South {
			directions = []direction{current.dir}
		} else {
			directions = []direction{North, South}
		}
	case '-':
		if current.dir == East || current.dir == West {
			directions = []direction{current.dir}
		} else {
			directions = []direction{East, West}
		}
	case '/':
		switch current.dir {
		case North:
			directions = []direction{East}
		case South:
			directions = []direction{West}
		case East:
			directions = []direction{North}
		case West:
			directions = []direction{South}
		}
	case '\\':
		switch current.dir {
		case North:
			directions = []direction{West}
		case South:
			directions = []direction{East}
		case East:
			directions = []direction{South}
		case West:
			directions = []direction{North}
		}
	}

	results := []beam{}
	for _, dir := range directions {
		results = append(results, beam{nextX, nextY, dir})
	}

	return results
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, b := range row {
			fmt.Printf("%c", b)
		}
		fmt.Print("\n")
	}
}
