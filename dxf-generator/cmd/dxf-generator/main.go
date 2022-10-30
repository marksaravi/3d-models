package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/marksaravi/dxfgenerator/epitrochoid"
)

const svg = `
<svg viewBox="%f %f %f %f" xmlns="http://www.w3.org/2000/svg">
 <style>svg { background-color: white; }</style>
 %s
</svg>
`

func main() {
	fmt.Println("DXF Generator")
	const scale float64 = 100
	const endAngle float64 = math.Pi
	const chunkSize = 180
	const numChunks = 10

	const da = endAngle / (chunkSize * numChunks)
	const R = float64(3) * scale
	const r = float64(1) * scale
	const d = float64(2) * scale
	var minx, miny float64 = 1000000, 1000000
	var maxx, maxy float64 = -1000000, -1000000
	polylines := strings.Builder{}

	for i := 0; i < numChunks; i++ {
		builder := strings.Builder{}
		comma := ""
		builder.WriteString("<polyline  stroke-width=\"1\" fill=\"none\" stroke=\"black\" points=\"")
		for j := 0; j < chunkSize; j++ {
			angle := float64(i*chunkSize+j) * da
			x, y := epitrochoid.Hypocycloid(angle, R, r, d)
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
	}

	fmt.Println(minx, miny, maxx, maxy)
	fout, _ := os.Create("./rotor.svg")
	fout.WriteString(fmt.Sprintf(svg, minx-50, miny-50, maxx+50, maxy+50, polylines.String()))
	fout.Close()
}
