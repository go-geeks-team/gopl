// Exercise 3.6

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	iterations = 100
	width      = 1024
	height     = 1024

	xmin, ymin, xmax, ymax = -2, -2, +2, +2

	red   = 200
	green = 200
	blue  = 200
)

func main() {
	rect := image.Rect(0, 0, width, height)
	img := image.NewNRGBA(rect)
	draw.Draw(img, rect, image.NewUniform(color.Black), image.Point{}, draw.Src)
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			y := float64(py)/float64(height)*(ymax-ymin) + ymin
			z := complex(x, y)
			result := mandelbrot(z)

			img.Set(px, py, color.NRGBA{
				R: uint8(red * result),
				G: uint8(green * result),
				B: uint8(blue * result),
				A: 255,
			})
		}
	}
	png.Encode(os.Stdout, img) // Примечание: игнорируем ошибки
}

func mandelbrot(xy complex128) float64 {
	cnt := 0
	for z := xy; cmplx.Abs(z) < 2 && cnt < iterations; cnt++ {
		z = z*z + xy
	}
	return float64(iterations-cnt) / iterations
}
