package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	bricks := []brick{}
	for _, brickDef := range strings.Split(inputString, "\n") {
		brickDef = strings.Replace(brickDef, "~", ",", -1)

		coordStrings := strings.Split(brickDef, ",")
		coords := []int{}
		for _, coordString := range coordStrings {
			coord, _ := strconv.Atoi(coordString)
			coords = append(coords, coord)
		}

		bricks = append(bricks, brick{
			point{coords[0], coords[1], coords[2]},
			point{coords[3], coords[4], coords[5]},
		})
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start.z < bricks[j].start.z
	})

	for i := range bricks {
		brick := bricks[i]
		z := 1

		prevIdx := i - 1
		for prevIdx >= 0 {
			prevBrick := bricks[prevIdx]

			if max(brick.start.x, prevBrick.start.x) <= min(brick.end.x, prevBrick.end.x) &&
				max(brick.start.y, prevBrick.start.y) <= min(brick.end.y, prevBrick.end.y) {
				z = max(prevBrick.end.z+1, z)
			}

			prevIdx -= 1
		}

		brick.end.z -= brick.start.z - z
		brick.start.z = z

		bricks[i] = brick
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start.z < bricks[j].start.z
	})

	supporting := map[brick]map[brick]struct{}{}
	supportedBy := map[brick]map[brick]struct{}{}

	for i := range bricks {
		current := bricks[i]

		for j := i + 1; j < len(bricks); j++ {
			compare := bricks[j]

			if current.end.z+1 < compare.start.z {
				break
			}

			if max(current.start.x, compare.start.x) <= min(current.end.x, compare.end.x) &&
				max(current.start.y, compare.start.y) <= min(current.end.y, compare.end.y) &&
				current.end.z+1 == compare.start.z {
				currentSupporting, ok := supporting[current]
				if !ok {
					currentSupporting = map[brick]struct{}{}
					supporting[current] = currentSupporting
				}
				currentSupporting[compare] = struct{}{}

				compareSupportedBy, ok := supportedBy[compare]
				if !ok {
					compareSupportedBy = map[brick]struct{}{}
					supportedBy[compare] = compareSupportedBy
				}
				compareSupportedBy[current] = struct{}{}
			}
		}
	}

	count := 0
	for _, b := range bricks {
		fallen := map[brick]struct{}{}

		queue := []brick{b}
		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]

			fallen[current] = struct{}{}

			for supportedByCurrent := range supporting[current] {
				supported := true
				for supporter := range supportedBy[supportedByCurrent] {
					if _, ok := fallen[supporter]; !ok {
						supported = false
						break
					}
				}
				if !supported {
					continue
				}
				queue = append(queue, supportedByCurrent)
			}
		}

		count += len(fallen) - 1
	}

	println(count)
}

type brick struct {
	start, end point
}

type point struct {
	x, y, z int
}
