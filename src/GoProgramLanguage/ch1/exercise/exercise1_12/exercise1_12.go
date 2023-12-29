package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu3 sync.Mutex
var count3 int

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/count", counter3)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		var cycles string
		for k, v := range r.Form {
			if k == "cycles" {
				cycles = v[0]
			}
		}
		cycle, _ := strconv.Atoi(cycles)
		lissajous(w, cycle)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func counter3(w http.ResponseWriter, r *http.Request) {
	mu3.Lock()
	fmt.Fprintf(w, "Count %d\n", count3)
	mu3.Unlock()
}

func lissajous(out io.Writer, cycles int) {
	const (
		//cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	if cycles == 0 {
		cycles = 5
	}
	freq := rand.Float64()
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
