// Surface вычисляет SVG-представление трехмерного графика функции.

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // Раазмер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапазон осей (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в еденице x или y
	zscale        = height * 0.4        // Пикселей в еденице z
	angle         = math.Pi / 6         // Углы осей x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, inf := corner(i+1, j)
			if inf {
				return
			}
			bx, by, inf := corner(i, j)
			if inf {
				return
			}
			cx, cy, inf := corner(i, j+1)
			if inf {
				return
			}
			dx, dy, inf := corner(i+1, j+1)
			if inf {
				return
			}
			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Ищем угловую точку (x, y) ячейки (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x, y, z) на двумернуб канву SVG (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, math.IsInf(z, 0)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0.0)
	return math.Sin(r) / r
}
