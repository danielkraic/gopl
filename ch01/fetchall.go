package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("time: %.2f\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	size, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%s: size %d, time %.2f sec", url, size, time.Since(start).Seconds())
}
