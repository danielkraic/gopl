package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mtx sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCount)
	http.HandleFunc("/info", handlerInfo)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	count++
	mtx.Unlock()

	fmt.Fprintf(w, "URL.Path = %s\n", r.URL.Path)
}

func handlerCount(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	fmt.Fprintf(w, "%d", count)
	mtx.Unlock()
}

func handlerInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)

	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header %s: %s\n", k, v)
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Form error: %v\n", err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form %s: %s\n", k, v)
	}
}
