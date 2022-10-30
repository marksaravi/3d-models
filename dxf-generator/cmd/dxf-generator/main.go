package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const scale float64 = 100
const chunkSize = 180
const da = math.Pi / chunkSize

const svg = `
<svg viewBox="%f %f %f %f" xmlns="http://www.w3.org/2000/svg">
 <style>svg { background-color: white; }</style>
 %s
</svg>
`

func genHousing(fixedRaduis, R, endAngle float64) string {
	rotatingRadius := fixedRaduis / 2 * 3
	var minx, miny float64 = 1000000, 1000000
	var maxx, maxy float64 = -1000000, -1000000
	polylines := strings.Builder{}
	e := rotatingRadius - fixedRaduis

	var angle float64 = 0
	for angle < endAngle {
		builder := strings.Builder{}
		comma := ""
		builder.WriteString("<polyline  stroke-width=\"1\" fill=\"none\" stroke=\"black\" points=\"")
		for j := 0; j < chunkSize; j++ {
			angle += da
			x := e*math.Cos(angle) + R*math.Cos(angle/3)
			y := e*math.Sin(angle) + R*math.Sin(angle/3)
			x *= scale
			y *= scale
			builder.WriteString(fmt.Sprintf("%s%f,%f", comma, x, y))
			comma = ","
			if x > maxx {
				maxx = x
			}
			if y > maxy {
				maxy = y
			}
			if x < minx {
				minx = x
			}
			if y < miny {
				miny = y
			}
		}
		builder.WriteString("\" />")
		polylines.WriteString(builder.String())
		fmt.Println(angle / math.Pi * 180)
	}

	fmt.Println(minx, miny, maxx, maxy)
	marginx := (maxx - minx) / 10
	marginy := (maxy - miny) / 10
	return fmt.Sprintf(svg, minx-marginx, miny-marginx, maxx-minx+marginx, maxy-miny+marginy, polylines.String())
}

func saveSvg(name, svg string) {
	fout, _ := os.Create(fmt.Sprintf("./%s.svg", name))
	fout.WriteString(svg)
	fout.Close()
}

func main() {
	saveSvg("housing", genHousing(24/10, 52/10, math.Pi*2))
	// saveSvg("rotor", 3, 2, 5, math.Pi*5, epitrochoid.Hypotrochoid)
}
