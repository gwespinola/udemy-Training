// Modify the lissajous program to produce images in multiple colours by adding more values to palette and then
// displaying them by changing the third argument of SetColorIndex in some interesting way.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x00,0x00,0x00, 0xff},
	color.RGBA{0x00,0xFF,0xFF, 0xff},
}

const (
	blackIndex = 	0	// Next colour in palette
	cyanIndex =  	1	// First colour in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles	= 5	// number of complete x oscillator revolutions
		res	= 0.001 // angular resolution
		size	= 100	// image canvas covers [-size..+size]
		nframes = 64	// number of animation frames
		delay	= 8	// delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0	// Relative frequency of y oscillator
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0		// Phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.2), size+int(y*size+0.8), cyanIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	//	NOTE: Ignoring encoding errors
}