// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions.

package main

import (
	"fmt"
	"time"

	"gopl/ch2/popcount"

	popcountrange "gopl/ch2/popcount2.3/popcount"
)

func main() {
	start := time.Now()
	for i := 0; i < 1e5; i++ {
		popcount.PopCount(uint64(i))
	}
	fmt.Printf("%s elapsed in population count by single expression\n", time.Since(start).String())

	start = time.Now()
	for i := 0; i < 1e5; i++ {
		popcountrange.PopCount(uint64(i))
	}
	fmt.Printf("%s elapsed in population count by range\n", time.Since(start).String())
}
