// Exercise 1.3:
// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		arg = " "
	}
	fmt.Printf("%s elapsed in concatenation\n", time.Since(start).String())

	start = time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Printf("%s elapsed in strings.Join\n", time.Since(start).String())
}
