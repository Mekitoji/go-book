// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	// roots of this function: 1+0i, 1+0i, 0+1i, 0-1i
	// use builin functions real() and imag() to find out the color of root
	var root int
	if real(z) > 0 {
		if imag(z) > 0 {
			root = 1
		} else {
			root = 2
		}
	} else {
		if imag(z) > 0 {
			root = 3
		} else {
			root = 4
		}
	}

	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		var clr = 255 - contrast*i
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			switch root {
			case 1:
				return color.RGBA{clr, 0, 0, 255}
			case 2:
				return color.RGBA{0, clr, 0, 255}
			case 3:
				return color.RGBA{0, 0, clr, 255}
			case 4:
				return color.RGBA{clr, 0, clr, 255}
			}
		}
	}
	return color.Black
}
