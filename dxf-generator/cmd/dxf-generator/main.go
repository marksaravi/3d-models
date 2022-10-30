package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/marksaravi/dxfgenerator/epitrochoid"
)

func main() {
	fmt.Println("DXF Generator")
	const scale float64 = 100
	const da = math.Pi / (180 * 5)
	const R = float64(3) * scale
	const r = float64(1) * scale
	const d = float64(2) * scale
	builder := strings.Builder{}
	var min float64 = 0
	var max float64 = 0
	var offset float64 = 1000
	for angle := float64(0); angle < 2*math.Pi; angle += da {
		x, y := epitrochoid.Hypocycloid(angle, R, r, d)
		x += offset
		y += offset
		builder.WriteString(fmt.Sprintf(",%f,%f", x, y))
		if x > max {
			max = x
		}
		if y > max {
			max = y
		}
		if x < min {
			min = x
		}
		if y < min {
			min = y
		}
	}
	fmt.Println(min, max)
	fout, _ := os.Create("./data")
	fout.WriteString(builder.String())
	fout.Close()
}
