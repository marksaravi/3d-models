package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const chunkSize = 180

const svg = `
<svg viewBox="%f %f %f %f" xmlns="http://www.w3.org/2000/svg">
 <style>svg { background-color: white; }</style>
 %s
</svg>
`

func housing(angle, e, R float64) (x, y float64) {
	x = e*math.Cos(angle) + R*math.Cos(angle/3)
	y = e*math.Sin(angle) + R*math.Sin(angle/3)
	return
}

// func rotor(angle, e, R float64) (x, y float64) {
// 	x = R*math.Cos(2*angle) + 3*e*e/2/R*(math.Cos(8*angle)+math.Cos(4*angle)) - e*math.Sqrt(1-9*e*e/R/R*math.Sin(3*angle)*math.Sin(3*angle)*(math.Cos(5*angle)+math.Cos(angle)))
// 	y = R*math.Sin(2*angle) + 3*e*e/2/R*(math.Sin(8*angle)+math.Sin(4*angle)) + e*math.Sqrt(1-9*e*e/R/R*math.Sin(3*angle)*math.Sin(3*angle)*(math.Sin(5*angle)+math.Sin(angle)))
// 	return
// }

func genHousing(fixedRaduis, R, startAngle, endAngle float64, fn func(angle, e, R float64) (x, y float64)) string {
	fmt.Println(startAngle/math.Pi*180, endAngle/math.Pi*180)
	rotatingRadius := fixedRaduis / 2 * 3
	var minx, miny float64 = 1000000, 1000000
	var maxx, maxy float64 = -1000000, -1000000
	polylines := strings.Builder{}
	e := rotatingRadius - fixedRaduis

	da := (endAngle - startAngle) / chunkSize
	var angle float64 = startAngle
	for angle < endAngle {
		builder := strings.Builder{}
		comma := ""
		builder.WriteString("<polyline  stroke-width=\"1\" fill=\"none\" stroke=\"black\" points=\"")
		for j := 0; j < chunkSize; j++ {
			angle += da
			x, y := fn(angle, e, R)
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
	width := maxx - minx
	height := maxy - miny
	xc := minx + width/2
	yc := miny + height/2
	polylines.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" stroke=\"black\" />", minx+50, miny, minx+50+R, miny))
	polylines.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" stroke-width=\"1\" fill=\"none\" stroke=\"black\" />", xc+550, yc-250, rotatingRadius))
	polylines.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" stroke-width=\"1\" fill=\"none\" stroke=\"black\" />", xc+550, yc-250, fixedRaduis))

	fmt.Println(minx, miny, maxx, maxy)
	marginx := width / 10
	marginy := height / 10
	return fmt.Sprintf(svg, minx-marginx, miny-marginx, width+marginx, height+marginy, polylines.String())
}

func saveSvg(name, svg string) {
	fout, _ := os.Create(fmt.Sprintf("./%s.svg", name))
	fout.WriteString(svg)
	fout.Close()
}

func main() {
	saveSvg("housing", genHousing(380, 1260, 0, math.Pi*3, housing))
	// saveSvg("rotor", genHousing(38/10, 126/10, math.Pi/6, math.Pi/2, rotor))
	// saveSvg("rotor", 3, 2, 5, math.Pi*5, epitrochoid.Hypotrochoid)
}
