package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		samples                = 4
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	n := math.Sqrt(samples)
	for py := 0; py < height; py++ {
		y1 := float64(py*2)/(height*n)*(ymax-ymin) + ymin
		y2 := float64(2*py+1)/(height*n)*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x1 := float64(px*2)/(width*n)*(xmax-xmin) + xmin
			x2 := float64(2*px+1)/(width*n)*(xmax-xmin) + xmin

			z1 := mandelbrot(complex(x1, y1)).Y
			z2 := mandelbrot(complex(x2, y1)).Y
			z3 := mandelbrot(complex(x1, y2)).Y
			z4 := mandelbrot(complex(x2, y2)).Y
			// Image point (px, py) represents complex value z.
			z := (z1 + z2 + z3 + z4) / samples
			img.Set(px, py, color.Gray{z})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Gray {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Gray{0}
}
