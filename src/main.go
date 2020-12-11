package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	width        int = 2000
	height       int = 1000
	maxIteration int = 180
)

func scale(in, inMin, inMax int, outMin, outMax float64) float64 {
	var n float64 = float64(in-inMin) / float64(inMax-inMin)
	var out float64 = n*(outMax-outMin) + outMin
	return out
}
func fractal(img *image.NRGBA, ale float64) {

	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			cx, cy := scale(px, 0, width, -2.511, 1), scale(py, 0, height, -1, 1)
			var c = complex(cx, cy)
			var z complex128 = 0 + 0i
			var i int = 0
			for cmplx.Abs(cmplx.Conj(z)) < 2*4 && i < maxIteration {
				z = (cmplx.Sqrt(cmplx.Sinh((z*z*z*z)*math.E)/math.Pi) + cmplx.Sin(c)) * complex(ale, 0)
				i++
			}
			img.Set(px, py, color.NRGBA{
				R: uint8(i % 64 * 64), G: uint8(i % 8 * 32), B: uint8(i % 64 * 64), A: 255,
			})
		}
	}
}

func main() {
	var img *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, width, height))
	var img2 *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, width, height))
	var img3 *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, width, height))
	fractal(img, 1)
	f1, _ := os.Create("../image/fractalx1.png")
	defer f1.Close()
	png.Encode(f1, img)
	f2, _ := os.Create("../image/fractalx2.png")
	defer f2.Close()
	fractal(img2, 2)
	png.Encode(f2, img2)
	f3, _ := os.Create("../image/fractalx3.png")
	defer f3.Close()
	fractal(img3, 3)
	png.Encode(f3, img3)
}
