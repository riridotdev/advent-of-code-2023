package main

import (
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	moduleDefs := strings.Split(strings.TrimSpace(string(inputBytes)), "\n")

	inputs := map[string][]string{}
	for _, def := range moduleDefs {
		parts := strings.Split(def, "->")

		module := strings.TrimSpace(parts[0])
		outputsString := strings.TrimSpace(parts[1])

		var name string
		if module[0] == '%' || module[0] == '&' {
			name = module[1:]
		} else {
			name = module
		}

		for _, output := range strings.Split(outputsString, ",") {
			output = strings.TrimSpace(output)
			inputsForOutput, ok := inputs[output]
			if !ok {
				inputsForOutput = []string{}
			}
			inputs[output] = append(inputsForOutput, name)
		}
	}

	modules := map[string]module{}

	for _, def := range moduleDefs {
		parts := strings.Split(def, "->")

		module := strings.TrimSpace(parts[0])
		outputsString := strings.TrimSpace(parts[1])

		outputs := []string{}
		for _, output := range strings.Split(outputsString, ",") {
			output = strings.TrimSpace(output)
			outputs = append(outputs, output)
		}

		name := module[1:]
		switch module[0] {
		case '%':
			modules[name] = &flipflop{low, outputs}
		case '&':
			modules[name] = newConjunction(inputs[name], outputs)
		case 'b':
			modules["broadcaster"] = broadcaster(outputs)
		}
	}

	lowPulses := 0
	highPulses := 0

	type signalStep struct {
		signal
		source string
	}

	for i := 0; i < 1000; i++ {
		queue := []signalStep{{signal{"broadcaster", low}, ""}}
		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]

			if current.state == low {
				lowPulses += 1
			}
			if current.state == high {
				highPulses += 1
			}

			if current.target == "output" {
				continue
			}

			module, ok := modules[current.target]
			if !ok {
				continue
			}

			for _, signal := range module.process(current.source, current.state) {
				queue = append(queue, signalStep{signal, current.target})
			}

		}
	}

	println(lowPulses * highPulses)
}

type module interface {
	process(source string, s state) []signal
}

type broadcaster []string

func (b broadcaster) process(_ string, s state) []signal {
	signals := []signal{}
	for _, output := range b {
		signals = append(signals, signal{output, s})
	}
	return signals
}

type flipflop struct {
	state   state
	outputs []string
}

func (f *flipflop) process(_ string, s state) []signal {
	if s == high {
		return nil
	}

	f.state = !f.state

	signals := []signal{}
	for _, output := range f.outputs {
		signals = append(signals, signal{output, f.state})
	}

	return signals
}

type conjunction struct {
	inputs  map[string]state
	outputs []string
}

func newConjunction(inputs []string, outputs []string) *conjunction {
	c := &conjunction{map[string]state{}, outputs}
	for _, input := range inputs {
		c.inputs[input] = low
	}
	return c
}

func (c *conjunction) process(input string, s state) []signal {
	c.inputs[input] = s

	state := low
	for _, v := range c.inputs {
		if v == low {
			state = high
			break
		}
	}

	signals := []signal{}
	for _, output := range c.outputs {
		signals = append(signals, signal{output, state})
	}

	return signals
}

type signal struct {
	target string
	state  state
}

type state bool

const (
	low  state = false
	high state = true
)
