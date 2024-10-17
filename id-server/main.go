package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var id atomic.Uint64

func idHandler(w http.ResponseWriter, req *http.Request) {
	// Can use here, but then have to make call to id.Load()
	// id.Add(1)

	fmt.Fprintf(w, "%d", id.Add(1))
}

func main() {
	http.HandleFunc("/serve-id", idHandler)
	http.ListenAndServe(":8090", nil)
}
