package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

type node struct {
	point
	cost      int
	direction point
	steps     int
}

type memo struct {
	point
	direction point
	steps     int
}

const minSteps = 4
const maxSteps = 10

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

	visited := map[memo]struct{}{}

	pq := priorityQueue{[]node{}}
	pq.insert(node{point{0, 0}, 0, point{0, 0}, math.MaxInt})
	for {
		current := pq.queue[0]
		pq.queue = pq.queue[1:]

		if current.x == len(grid[0])-1 && current.y == len(grid)-1 && current.steps >= minSteps {
			println(current.cost)
			break
		}

		if _, ok := visited[memo{current.point, current.direction, current.steps}]; ok {
			continue
		}
		visited[memo{current.point, current.direction, current.steps}] = struct{}{}

		directions := [4]point{
			{0, 1},
			{1, 0},
			{-1, 0},
			{0, -1},
		}
		for _, direction := range directions {
			nextX, nextY := current.x+direction.x, current.y+direction.y
			if nextX >= len(grid[0]) || nextX < 0 || nextY >= len(grid) || nextY < 0 {
				continue
			}
			if current.direction.x == -direction.x && current.direction.y == -direction.y {
				continue
			}
			if current.direction != direction {
				if current.steps < minSteps {
					continue
				}
				pq.insert(node{point{nextX, nextY}, current.cost + must(strconv.Atoi(string(grid[nextY][nextX]))), direction, 1})
				continue
			}
			if current.steps == maxSteps {
				continue
			}
			pq.insert(node{point{nextX, nextY}, current.cost + must(strconv.Atoi(string(grid[nextY][nextX]))), direction, current.steps + 1})
		}
	}
}

type priorityQueue struct {
	queue []node
}

func (pq *priorityQueue) insert(n node) {
	insertIdx := -1
	for i := range pq.queue {
		if n.cost > pq.queue[i].cost {
			continue
		}
		insertIdx = i
		break
	}

	if insertIdx == -1 {
		pq.queue = append(pq.queue, n)
		return
	}

	pq.queue = append(pq.queue, pq.queue[len(pq.queue)-1])
	for i := len(pq.queue) - 1; i > insertIdx; i-- {
		pq.queue[i] = pq.queue[i-1]
	}

	pq.queue[insertIdx] = n
}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
