package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(inputBytes))

	sections := strings.Split(inputString, "\n\n")

	directions := sections[0]
	networkMap := sections[1]

	nodes := strings.Split(networkMap, "\n")

	startNodes := []string{}
	network := map[string][2]string{}
	for _, node := range nodes {
		nodeParts := strings.Split(node, "=")

		name := strings.TrimSpace(nodeParts[0])
		paths := strings.TrimSpace(nodeParts[1])

		if name[2] == 'A' {
			startNodes = append(startNodes, name)
		}

		network[name] = [2]string{paths[1:4], paths[6:9]}
	}

	cycles := []int{}
	for _, node := range startNodes {
		count := 0
		firstZ := -1

	outer:
		for {
			for _, direction := range directions {
				count += 1
				nodePaths := network[node]

				if direction == 'L' {
					node = nodePaths[0]
				}
				if direction == 'R' {
					node = nodePaths[1]
				}

				if node[2] != 'Z' {
					continue
				}
				if firstZ == -1 {
					firstZ = count
					continue
				}

				cycles = append(cycles, count-firstZ)
				break outer
			}
		}
	}

	lcm := 1
	for _, cycle := range cycles {
		lcm = lowestCommonMultiple(lcm, cycle)
	}

	fmt.Println(lcm)
}

func lowestCommonMultiple(i, j int) int {
	return (i * j) / largestCommonDivisor(i, j)
}

func largestCommonDivisor(i, j int) int {
	k := min(i, j)
	for {
		if i%k == 0 && j%k == 0 {
			return k
		}
		k -= 1
	}
}
