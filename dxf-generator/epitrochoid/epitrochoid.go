package epitrochoid

import "math"

func Epitrochoid(angle, R, r, d float64) (x, y float64) {
	x = (R+r)*math.Cos(angle) - d*math.Cos((R+r)/r*angle)
	y = (R+r)*math.Sin(angle) - d*math.Sin((R+r)/r*angle)
	return
}
