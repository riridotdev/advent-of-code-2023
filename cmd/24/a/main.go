package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("input")
	input := strings.Replace(string(inputBytes), " @", ",", -1)

	hailstones := []hailstone{}
	for _, def := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(def, ",")

		hailstone := hailstone{}

		setters := []func(float64){
			func(f float64) { hailstone.x = f },
			func(f float64) { hailstone.y = f },
			func(f float64) { hailstone.z = f },
			func(f float64) { hailstone.vX = f },
			func(f float64) { hailstone.vY = f },
			func(f float64) { hailstone.vZ = f },
		}
		for i, setter := range setters {
			val, err := strconv.Atoi(strings.TrimSpace(parts[i]))
			if err != nil {
				panic(err)
			}
			setter(float64(val))
		}

		hailstones = append(hailstones, hailstone)
	}

	count := 0
	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {
			a1, b1, c1 := h1.vY, -h1.vX, (h1.x*h1.vY)-(h1.y*h1.vX)
			a2, b2, c2 := h2.vY, -h2.vX, (h2.x*h2.vY)-(h2.y*h2.vX)

			if a1*b2 == b1*a2 {
				continue
			}

			x := ((c1 * b2) - (c2 * b1)) / ((a1 * b2) - (a2 * b1))
			y := ((c2 * a1) - (c1 * a2)) / ((a1 * b2) - (a2 * b1))

			const rangeMin = 200000000000000
			const rangeMax = 400000000000000
			if x > rangeMax || x < rangeMin {
				continue
			}
			if y > rangeMax || y < rangeMin {
				continue
			}

			valid := true
			for _, h := range []hailstone{h1, h2} {
				if (x-h.x)*h.vX < 0 {
					valid = false
					break
				}
				if (y-h.y)*h.vY < 0 {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			count += 1
		}
	}

	println(count)
}

type hailstone struct {
	x, y, z, vX, vY, vZ float64
}
