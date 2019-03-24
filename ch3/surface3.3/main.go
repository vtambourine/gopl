package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var zmin, zmax = math.Inf(1), math.Inf(-1)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			if !math.IsNaN(az) {
				if az > zmax {
					zmax = az
				}
				if az < zmin {
					zmin = az
				}
			}

			if !math.IsNaN(bz) {
				if bz > zmax {
					zmax = bz
				}
				if bz < zmin {
					zmin = bz
				}
			}

			if !math.IsNaN(cz) {
				if cz > zmax {
					zmax = cz
				}
				if cz < zmin {
					zmin = cz
				}
			}

			if !math.IsNaN(dz) {
				if dz > zmax {
					zmax = dz
				}
				if dz < zmin {
					zmin = dz
				}
			}

			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Printf("<polygon fill='%s' stroke='%s' "+
				"points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j), "#fff",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func color(i, j int) string {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	red := (z - zmin) / (zmax - zmin) * 255
	// blue := 255 - (z-zmin)/(zmax-zmin)*255

	return fmt.Sprintf("#%02x00%02x", int(red), int(0))
}
