package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/marksaravi/dxfgenerator/epitrochoid"
)

const scale float64 = 100
const endAngle float64 = math.Pi * 2
const chunkSize = 180
const numChunks = 10
const da = endAngle / (chunkSize * numChunks)

const svg = `
<svg viewBox="%f %f %f %f" xmlns="http://www.w3.org/2000/svg">
 <style>svg { background-color: white; }</style>
 %s
</svg>
`

func genShape(R, r, d float64, fn func(angle, R, r, d float64) (x, y float64)) {
	var minx, miny float64 = 1000000, 1000000
	var maxx, maxy float64 = -1000000, -1000000
	polylines := strings.Builder{}

	for i := 0; i < numChunks; i++ {
		builder := strings.Builder{}
		comma := ""
		var angle float64
		builder.WriteString("<polyline  stroke-width=\"1\" fill=\"none\" stroke=\"black\" points=\"")
		for j := 0; j < chunkSize; j++ {
			angle = float64(i*chunkSize+j) * da
			x, y := fn(angle, R, r, d)
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
	fout, _ := os.Create("./housing.svg")
	marginx := (maxx - minx) / 10
	marginy := (maxy - miny) / 10
	fout.WriteString(fmt.Sprintf(svg, minx-marginx, miny-marginx, maxx-minx+marginx, maxy-miny+marginy, polylines.String()))
	fout.Close()
}

func main() {
	genShape(2, 1, 0.5, epitrochoid.Epitrochoid)
}
