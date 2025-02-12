// Exercise 3.9

package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func Gen(w io.Writer, r *http.Request) {
	var y float64
	var zoom = 1

	if r.URL.Query().Has("z") {
		zS := r.URL.Query().Get("z")
		a, _ := strconv.Atoi(zS)
		zoom = a
	}

	img := image.NewRGBA(image.Rect(0, 0, width*zoom, height*zoom))
	for py := 0; py < height*zoom; py++ {
		if r.URL.Query().Has("y") {
			a := r.URL.Query().Get("y")
			i, _ := strconv.ParseFloat(a, 64)
			y = float64(py)/height*(ymax-ymin) + ymin + i
		} else {
			y = float64(py)/height*(ymax-ymin) + ymin
		}

		for px := 0; px < width*zoom; px++ {
			var x float64
			if r.URL.Query().Has("x") {
				a := r.URL.Query().Get("x")
				i, _ := strconv.ParseFloat(a, 64)
				x = float64(px)/width*(xmax-xmin) + xmin + i
			} else {
				x = float64(px)/width*(xmax-xmin) + xmin
			}

			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px*zoom, py*zoom, mandelbrot(z))
		}
	}
	png.Encode(w, img) // Примечание: игнорируем ошибки
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 50
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: 128 - contrast*n, G: 255 - contrast*n, B: 77 - contrast*n, A: 255}
		}
	}
	return color.Black
}
