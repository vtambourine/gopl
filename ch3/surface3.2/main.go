package main

import (
	"fmt"
	"math"
	"os"
)

type surface func(float64, float64) float64

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	surfaces := map[string]surface{
		"drop":   drop,
		"eggbox": eggbox,
		"moguls": moguls,
		"saddle": saddle,
	}
	for n, s := range surfaces {
		file, err := os.Create(fmt.Sprintf("%s.svg", n))
		if err != nil {
			fmt.Fprintf(os.Stderr, "surface: %v\n", err)
			os.Exit(1)
		}
		svg(s, file)
		file.Close()
	}
}

func svg(f surface, out *os.File) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, f surface) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func drop(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return math.Sin(x) * math.Cos(y) * 0.2
}

func moguls(x, y float64) float64 {
	return math.Sin(x) * 0.2
}

func saddle(x, y float64) float64 {
	return y*y/225.0 - x*x/50.0
}
