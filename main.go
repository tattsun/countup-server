package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var (
	count uint64
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		atomic.AddUint64(&count, 1)
		cnt := atomic.LoadUint64(&count)
		fmt.Fprintf(w, "%d", cnt)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
