// Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time. Compare its performance to the table-lookup version.

package main

import (
	"fmt"
	"time"

	"gopl/ch2/popcount"

	popcountshift "gopl/ch2/popcount2.4/popcount"
)

func main() {
	start := time.Now()
	for i := 0; i < 1e5; i++ {
		popcount.PopCount(uint64(i))
	}
	fmt.Printf("%s elapsed in population count by single expression\n", time.Since(start).String())

	start = time.Now()
	for i := 0; i < 1e5; i++ {
		popcountshift.PopCount(uint64(i))
	}
	fmt.Printf("%s elapsed in population count by shift value\n", time.Since(start).String())
}
