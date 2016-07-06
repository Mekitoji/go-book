package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	red           = "#ff0000"
	blue          = "#0000ff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin 30, cos 30

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, isPeak := corner(i+1, j+1)
			var color string
			if isPeak {
				color = red
			} else {
				color = blue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	var isPeak bool
	if z > 0 {
		isPeak = true
	} else {
		isPeak = false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, isPeak
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if math.IsInf(math.Sin(r)/r, 0) {
		return 0
	}
	return math.Sin(r) / r
}
