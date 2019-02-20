// Exercise 1.4: Modify dup2 to print the names of all files in which
// each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	sources := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, sources)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup1.4: %v\n", err)
				continue
			}
			countLines(f, counts, sources)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for f := range sources[line] {
				fmt.Printf("%s ", f)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]int, sources map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		_, ok := sources[line]
		if !ok {
			sources[line] = make(map[string]int)
		}
		sources[line][f.Name()]++
	}
}
