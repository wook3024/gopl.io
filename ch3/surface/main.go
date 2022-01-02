// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/surface", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	var width, height float64 = 600, 320 // canvas size in pixels
	var cells float64 = 100              // number of grid cells
	var xyrange = 30.0                   // axis ranges (-xyrange..+xyrange)
	var xyscale = width / 2 / xyrange    // pixels per x or y unit
	var zscale = height * 0.4            // pixels per z unit
	var angle = math.Pi / 6              // angle of x, y axes (:=30°)
	var strokeColor = "grey"

	var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

	if r.FormValue("width") != "" {
		width, _ = strconv.ParseFloat(r.FormValue("width"), 64)
	}
	if r.FormValue("height") != "" {
		height, _ = strconv.ParseFloat(r.FormValue("height"), 64)
	}
	if r.FormValue("strokeColor") != "" {
		strokeColor = r.FormValue("strokeColor")
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%f' height='%f'>", strokeColor, width, height)
	fmt.Fprintf(w, "	<defs>"+
		"	<linearGradient id='grad' x2='0%%' y2='100%%'>"+
		"		<stop offset='0' stop-color='#ff0000'/>"+
		"		<stop offset='1' stop-color='#0000ff'/>"+
		"	</linearGradient>"+
		"	</defs>")
	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			ax, ay := corner(width, height, cos30, sin30, xyrange, xyscale, zscale, cells, i+1, j)
			bx, by := corner(width, height, cos30, sin30, xyrange, xyscale, zscale, cells, i, j)
			cx, cy := corner(width, height, cos30, sin30, xyrange, xyscale, zscale, cells, i, j+1)
			dx, dy := corner(width, height, cos30, sin30, xyrange, xyscale, zscale, cells, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='url(#grad)' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(width float64, height float64, cos30 float64, sin30 float64, xyrange float64, xyscale float64, zscale float64, cells float64, i, j int) (float64, float64) {
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

//!-
