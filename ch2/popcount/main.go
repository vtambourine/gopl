package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0.0; i < 1.0; i += 0.3 {
		x := math.Sin(i)
		y := math.Cos(i)
		fmt.Printf("%v %v %v\n", &i, &x, &y)
	}
}
