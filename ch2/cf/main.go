package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s, %s = %s\n",
			c, tempconv.CToF(c), f, tempconv.FToC(f))
	}
}
