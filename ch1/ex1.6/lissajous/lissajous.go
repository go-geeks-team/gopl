// Упражнение 1.6.
// Измените программу lissajous так, чтобы она генерировала изображения разных цветов, добавляя в палитру palette больше значений,
// а затем выводя их путём изменения третьего аргумента функции SetColorIndex некоторым нетривиальным образом.
//
// Lissajous генерирует анимированный GIF из случайных фигур Лисажу.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}, // black
	color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}, // red
	color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}, // green
	color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}, // blue
}

const (
	whiteIndex = 0 // Первый цвет палитры
	blackIndex = 1 // Следующий цвет палитры
)

func main() {
	createFiles()
}
func createFiles() bool {
	for i := 1; i < len(palette); i++ { // пропускаем элемент с индексом 0 (чёрный цвет)
		name := strconv.Itoa(i) + ".png"
		f, _ := os.Create(name)
		lissajous(f, uint8(i))
	}
	return true
}
func lissajous(out io.Writer, colorIndex uint8) {
	const (
		cycles  = 5     // x
		res     = 0.001 //
		size    = 100   // [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (еденица - 10мс)
	)

	// rand.Seed(time.Now().UTC().UnixNano()) // Deprecated
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}
