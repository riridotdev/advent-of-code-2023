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

	network := map[string][2]string{}
	for _, node := range nodes {
		nodeParts := strings.Split(node, "=")

		name := strings.TrimSpace(nodeParts[0])
		paths := strings.TrimSpace(nodeParts[1])

		network[name] = [2]string{paths[1:4], paths[6:9]}
	}

	steps := 0
	currentNode := "AAA"
	for {
		for _, direction := range directions {
			if currentNode == "ZZZ" {
				fmt.Println(steps)
				return
			}

			currentNodePaths := network[currentNode]

			if direction == 'L' {
				currentNode = currentNodePaths[0]
			}
			if direction == 'R' {
				currentNode = currentNodePaths[1]
			}

			steps += 1
		}
	}
}
