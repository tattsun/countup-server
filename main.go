package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

var (
	count uint64
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		atomic.AddUint64(&count, 1)
		cnt := atomic.LoadUint64(&count)

		hostname, err := os.Hostname()
		if err != nil {
			w.WriteHeader(500)
			fmt.Printf("err: %s", err)
			return
		}

		res := struct {
			Count    uint64
			Hostname string
		}{
			cnt,
			hostname,
		}

		enc := json.NewEncoder(w)
		enc.Encode(res)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
