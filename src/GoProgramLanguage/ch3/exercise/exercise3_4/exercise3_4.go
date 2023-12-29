// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	form := r.Form

	wd := 0
	wdQuery := form["width"]
	if len(wdQuery) > 0 {
		wdStr := wdQuery[0]
		wd, err = strconv.Atoi(wdStr)
		if err != nil {
			log.Print(err)
		}
	}

	ht := 0
	htQuery := form["height"]
	if len(htQuery) > 0 {
		htStr := htQuery[0]
		ht, err = strconv.Atoi(htStr)
		if err != nil {
			log.Print(err)
		}
	}

	stroke := ""
	strokeQuery := form["stroke"]
	if len(strokeQuery) > 0 {
		stroke = strokeQuery[0]
	}
	surface(w, wd, ht, stroke)
}

func surface(w http.ResponseWriter, wd, ht int, stroke string) {
	if wd <= 0 {
		wd = width
	}
	if ht <= 0 {
		ht = height
	}
	if len(stroke) == 0 {
		stroke = "grey"
	}
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", stroke, wd, ht)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
