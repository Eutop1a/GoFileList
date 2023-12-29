// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	form := r.Form
	var (
		Xmin, Xmax, Ymax, Ymin int
	)

	XmaxQ := form["xmax"]
	if len(XmaxQ) > 0 {
		XmaxStr := XmaxQ[0]
		Xmax, err = strconv.Atoi(XmaxStr)
		if err != nil {
			log.Print(err)
		}
	}

	if Xmax <= 0 {
		Xmax = xmax
	}

	XminQ := form["xmin"]
	if len(XminQ) > 0 {
		XminStr := XminQ[0]
		Xmin, err = strconv.Atoi(XminStr)
		if err != nil {
			log.Print(err)
		}
	}

	if Xmin <= 0 {
		Xmin = xmin
	}

	YmaxQ := form["ymax"]
	if len(YmaxQ) > 0 {
		YmaxStr := YmaxQ[0]
		Ymax, err = strconv.Atoi(YmaxStr)
		if err != nil {
			log.Print(err)
		}
	}

	if Ymax <= 0 {
		Ymax = ymax
	}

	YminQ := form["ymin"]
	if len(YminQ) > 0 {
		YminStr := YminQ[0]
		Ymin, err = strconv.Atoi(YminStr)
		if err != nil {
			log.Print(err)
		}
	}

	if Ymin <= 0 {
		Ymin = ymin
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(float64(Ymax)-float64(Ymin)) + float64(Ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(float64(Xmax)-float64(Xmin)) + float64(Xmin)
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
