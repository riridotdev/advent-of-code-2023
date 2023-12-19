package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	input := strings.Split((strings.TrimSpace(string(inputBytes))), "\n\n")

	pipelineDefs := input[0]
	partDefs := input[1]

	pipelines := map[string]pipeline{}
	for _, def := range strings.Split(pipelineDefs, "\n") {
		name, pipeline := parsePipeline(def)
		pipelines[name] = pipeline
	}

	total := 0
	for _, def := range strings.Split(partDefs, "\n") {
		c := parseComponent(def)
		result := "in"
		for result != "A" && result != "R" {
			result = pipelines[result].process(c)
		}

		if result == "A" {
			for _, v := range c {
				total += v
			}
		}
	}

	println(total)
}

type processor interface {
	process(component) string
}

type component map[byte]int

type pipeline []processor

func (p pipeline) process(comp component) string {
	for _, c := range p {
		if result := c.process(comp); result != "" {
			return result
		}
	}
	panic("invalid pipeline")
}

type conditional struct {
	target    byte
	condition condition
	amount    int
	onTrue    string
}

func (c conditional) process(comp component) string {
	if c.condition == gt {
		if comp[c.target] > c.amount {
			return c.onTrue
		}
		return ""
	}
	if c.condition == lt {
		if comp[c.target] < c.amount {
			return c.onTrue
		}
		return ""
	}
	panic("invalid condition type")
}

type pipelineResult string

func (pr pipelineResult) process(comp component) string {
	return string(pr)
}

type condition int

const (
	gt condition = iota
	lt
)

func parsePipeline(s string) (string, pipeline) {
	l, r := 0, 0

	p := pipeline{}

	for s[r] != '{' {
		r += 1
	}

	name := s[l:r]

	r += 1
	l = r

	stageDefs := strings.Split(s[l:len(s)-1], ",")
	for _, def := range stageDefs[:len(stageDefs)-1] {
		p = append(p, parseConditional(def))
	}

	p = append(p, pipelineResult(stageDefs[len(stageDefs)-1]))

	return name, p
}

func parseConditional(s string) conditional {
	c := conditional{}

	c.target = s[0]

	switch s[1] {
	case '>':
		c.condition = gt
	case '<':
		c.condition = lt
	default:
		panic("invalid condition")
	}

	r := 2
	for s[r] != ':' {
		r += 1
	}

	amount, err := strconv.Atoi(s[2:r])
	if err != nil {
		panic(err)
	}
	c.amount = amount

	r += 1

	c.onTrue = s[r:]

	return c
}

func parseComponent(s string) component {
	parts := strings.Split(s[1:len(s)-1], ",")
	c := component{}

	fields := []byte{'x', 'm', 'a', 's'}
	for i, field := range fields {
		val, err := strconv.Atoi(parts[i][2:])
		if err != nil {
			panic(err)
		}
		c[field] = val
	}

	return c
}
