// Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version of PopCount that counts bits by using this fact, and assess its performance.

package main

import (
	"fmt"
	"time"

	"gopl/ch2/popcount"

	popcountclearright "gopl/ch2/popcount2.4/popcount"
)

func main() {
	start := time.Now()
	for i := 0; i < 1e5; i++ {
		popcount.PopCount(uint64(i))
	}
	fmt.Printf("%s elapsed in population count by single expression\n", time.Since(start).String())

	start = time.Now()
	for i := 0; i < 1e5; i++ {
		popcountclearright.PopCount(uint64(i))
	}
	fmt.Printf("%s elapsed in population count by clearing rightmost\n", time.Since(start).String())
}
