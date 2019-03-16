package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl/ch2/cf2.2/unitconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			i := input.Text()
			conv(i)
		}
	} else {
		for _, arg := range os.Args[1:] {
			conv(arg)
		}
	}
}

func conv(s string) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf2.2: %v\n", err)
		os.Exit(1)
	}

	c := unitconv.Celsius(v)
	f := unitconv.Fahrenheit(v)
	m := unitconv.Meter(v)
	ft := unitconv.Foot(v)
	k := unitconv.Kilogram(v)
	p := unitconv.Pound(v)

	fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
		c, unitconv.CToF(c), f, unitconv.FToC(f),
		m, unitconv.MToF(m), ft, unitconv.FToM(ft),
		k, unitconv.KToP(k), p, unitconv.PToK(p))
}
