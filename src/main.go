package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	width        int = 3000
	height       int = 1500
	maxIteration int = 80
)

func scale(in, inMin, inMax int, outMin, outMax float64) float64 {
	var n float64 = float64(in-inMin) / float64(inMax-inMin)
	var out float64 = n*(outMax-outMin) + outMin
	return out
}
func fractal(img *image.NRGBA) {

	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			cx, cy := scale(px, 0, width, -2.511, 1), scale(py, 0, height, -1, 1)
			var c = complex(cx, cy)
			var z complex128 = 0 + 0i
			var i int = 0
			for cmplx.Abs(cmplx.Conj(z)) < 2*4 && i < maxIteration {
				z = (cmplx.Sqrt(cmplx.Sinh(z*z*z*z)) + c) * 2
				i++
			}
			img.Set(px, py, color.NRGBA{
				R: uint8(i % 16 * 16), G: uint8(i % 8 * 32), B: uint8(i % 4 * 64), A: 255,
			})
		}
	}
}

func main() {
	var img *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, width, height))
	fractal(img)
	f, _ := os.Create("../image/fractal.png")
	defer f.Close()
	png.Encode(f, img)

}
