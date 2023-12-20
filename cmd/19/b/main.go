package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	pipelineDefs := strings.Split((strings.TrimSpace(string(inputBytes))), "\n\n")[0]

	pipelines := map[string]pipeline{}
	for _, def := range strings.Split(pipelineDefs, "\n") {
		name, pipeline := parsePipeline(def)
		pipelines[name] = pipeline
	}

	r := componentRange{}
	fields := []byte{'x', 'm', 'a', 's'}
	for _, field := range fields {
		r[field] = fieldRange{1, 4000}
	}

	count := 0

	successRanges := []successRange{{r, "in"}}
	for len(successRanges) != 0 {
		current := successRanges[0]
		successRanges = successRanges[1:]

		if current.target == "A" {
			thisCount := 1
			for _, v := range current.componentRange {
				thisCount *= (v.max - v.min) + 1
			}
			count += thisCount
			continue
		}
		if current.target == "R" {
			continue
		}

		pipeline := pipelines[current.target]
		sRanges, _ := pipeline.findRanges(current.componentRange)

		successRanges = append(successRanges, sRanges...)
	}

	println(count)
}

type processor interface {
	findRanges(componentRange) ([]successRange, componentRange)
}

type pipeline []processor

func (p pipeline) findRanges(r componentRange) ([]successRange, componentRange) {
	successRanges := []successRange{}
	for _, c := range p {
		sRanges, fRange := c.findRanges(r)
		successRanges = append(successRanges, sRanges...)
		r = fRange
	}
	return successRanges, r
}

type conditional struct {
	target    byte
	condition condition
	amount    int
	onTrue    string
}

func (c conditional) findRanges(r componentRange) ([]successRange, componentRange) {
	sRange := componentRange{}
	fRange := componentRange{}
	for k, v := range r {
		sRange[k] = v
		fRange[k] = v
	}
	if c.condition == gt {
		sRange[c.target] = fieldRange{max(r[c.target].min, c.amount+1), r[c.target].max}
		fRange[c.target] = fieldRange{r[c.target].min, min(r[c.target].max, c.amount)}
	}
	if c.condition == lt {
		sRange[c.target] = fieldRange{r[c.target].min, min(r[c.target].max, c.amount-1)}
		fRange[c.target] = fieldRange{max(r[c.target].min, c.amount), r[c.target].max}
	}
	return []successRange{{sRange, c.onTrue}}, fRange
}

type pipelineResult string

func (pr pipelineResult) findRanges(r componentRange) ([]successRange, componentRange) {
	return []successRange{{r, string(pr)}}, nil
}

type successRange struct {
	componentRange componentRange
	target         string
}

type componentRange map[byte]fieldRange

type fieldRange struct {
	min, max int
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
