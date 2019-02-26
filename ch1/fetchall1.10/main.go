// Exercise 1.10

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.Create(fmt.Sprintf("fetchall-%d.out", time.Now().Unix()))
	if err != nil {
		fmt.Printf("fetchall: %v\n", err)
		return
	}
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		res := <-ch
		file.WriteString(res)
		fmt.Println(res)
	}
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
	file.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), nbytes, url)
}
