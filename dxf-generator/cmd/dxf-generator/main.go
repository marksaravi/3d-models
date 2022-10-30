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
	const da = math.Pi / 360
	const R = float64(3)
	const r = float64(1)
	const d = float64(0.5)
	var counter = 0
	builder := strings.Builder{}
	var prevx, prevy float64
	for angle := float64(0); angle < 2*math.Pi; angle += da {
		x, y := epitrochoid.Epitrochoid(angle, R, r, d)
		if counter == 0 {
			prevx = x
			prevy = y
			counter = 32
			continue
		}
		line := fmt.Sprintf(`
0
LINE
  5
%d
100
AcDbEntity
  8
0
100
AcDbLine
 10
%f
 20
%f
 30
0.0
 11
%f
 21
%f
 31
0.0
		`, counter, prevx, prevy, x, y)
		builder.WriteString(line)
		counter++
	}
	fout, _ := os.Create("./data.dxf")
	fout.WriteString(builder.String())
	fout.Close()
}
