// Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding more values to palette and then displaying them by changing the third argument of SetColorIndex in some interesting way.

package main

import (
	"image"
	"image/color/palette"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		delay   = 8
		nframes = 64
		res     = 0.001
		size    = 100
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	psize := len(palette.WebSafe)
	phase := 0.0
	for i := 0; i < nframes; i++ {
		c := 0.0
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette.WebSafe)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(int(c)%psize))
			c += 0.01
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
