package main

import (
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

	boundaryLength := 0

	position := point{0, 0}
	points := []point{position}
	for _, command := range strings.Split(inputString, "\n") {
		parts := strings.Fields(command)
		hex := parts[2][2 : len(parts[2])-1]

		direction := hex[5]
		vector := vectorForDirection(direction)

		distance := hex[:5]
		count, _ := strconv.ParseInt(distance, 16, 64)

		boundaryLength += int(count)

		position.x += vector.x * int(count)
		position.y += vector.y * int(count)

		points = append(points, point{position.x, position.y})
	}

	sum := 0
	for i := range points {
		aIdx := (i + 1) % len(points)
		bIdx := (len(points) + i - 1) % len(points)
		sum += points[i].x * (points[aIdx].y - points[bIdx].y)
	}

	area := abs(sum) / 2
	interior := area - (boundaryLength / 2) + 1

	println(interior + boundaryLength)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func vectorForDirection(dir byte) point {
	switch dir {
	case '3':
		return point{0, -1}
	case '1':
		return point{0, 1}
	case '2':
		return point{-1, 0}
	case '0':
		return point{1, 0}
	}
	panic("invalid direction")
}
