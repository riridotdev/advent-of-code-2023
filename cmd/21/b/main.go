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

	size := len(grid)

	const steps = 26501365

	gridReach := steps/size - 1

	oddGridCount := ((gridReach / 2) * 2) + 1
	oddGridCount *= oddGridCount

	evenGridCount := ((gridReach + 1) / 2) * 2
	evenGridCount *= evenGridCount

	oddPoints := countLocations(grid, start, size*2+1)
	evenPoints := countLocations(grid, start, size*2)

	corner_t := countLocations(grid, point{size - 1, start.y}, size-1)
	corner_r := countLocations(grid, point{start.x, 0}, size-1)
	corner_b := countLocations(grid, point{0, start.y}, size-1)
	corner_l := countLocations(grid, point{start.x, size - 1}, size-1)

	small_tr := countLocations(grid, point{size - 1, 0}, (size/2)-1)
	small_tl := countLocations(grid, point{size - 1, size - 1}, (size/2)-1)
	small_br := countLocations(grid, point{0, 0}, (size/2)-1)
	small_bl := countLocations(grid, point{0, size - 1}, (size/2)-1)

	large_tr := countLocations(grid, point{size - 1, 0}, ((size*3)/2)-1)
	large_tl := countLocations(grid, point{size - 1, size - 1}, ((size*3)/2)-1)
	large_br := countLocations(grid, point{0, 0}, ((size*3)/2)-1)
	large_bl := countLocations(grid, point{0, size - 1}, ((size*3)/2)-1)

	println(
		oddGridCount*oddPoints +
			evenGridCount*evenPoints +
			corner_t + corner_r + corner_b + corner_l +
			(gridReach+1)*(small_tr+small_tl+small_br+small_bl) +
			(gridReach)*(large_tr+large_tl+large_br+large_bl),
	)
}

func countLocations(grid [][]byte, start point, steps int) int {
	count := 0
	seen := map[point]struct{}{}
	queue := []point{start}
	for i := steps; i > 0; i-- {
		for j := len(queue); j > 0; j-- {
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
				next := point{nextX, nextY}
				if nextX >= len(grid[0]) || nextX < 0 || nextY >= len(grid) || nextY < 0 {
					continue
				}
				if grid[nextY][nextX] == '#' {
					continue
				}
				if _, ok := seen[next]; ok {
					continue
				}
				seen[next] = struct{}{}
				if i%2 == 1 {
					count += 1
				}
				queue = append(queue, next)
			}
		}
	}

	return count
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
