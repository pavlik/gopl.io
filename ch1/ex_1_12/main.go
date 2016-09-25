// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

func generateColorPalette(colorsCount int) []color.Color {
	// генерируем colorsCount случайных цветов и добавляем их в массив палитры
	palette := []color.Color{}

	for i := 0; i < colorsCount; i++ {
		random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
		r := uint8(random.Intn(255))
		g := uint8(random.Intn(255))
		b := uint8(random.Intn(255))
		palette = append(palette, color.RGBA{r, g, b, 1})
	}

	return palette
}

func getRandColorFromPalette(colorsCount int) uint8 {
	return uint8(rand.Intn(colorsCount))
}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.

	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", lissajous)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

func lissajous(w http.ResponseWriter, r *http.Request) {
	colorsCount := 10
	palette := generateColorPalette(colorsCount)

	var cycles = 5.0 // number of complete x oscillator revolutions := 5.0
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	if r.FormValue("cycles") != "" {
		c, err := strconv.Atoi(r.FormValue("cycles"))
		fmt.Println(c)
		if err != nil {
			log.Fatal(err)
		}

		cycles = float64(c)
	}

	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	currentColor := getRandColorFromPalette(colorsCount)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				currentColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim) // NOTE: ignoring encoding errors
}
