package main

import (
	"fmt"
	"net/http"
)

var id = 0

func idHandler(w http.ResponseWriter, req *http.Request) {
	id += 1
	fmt.Fprintf(w, "%d", id)
}

func main() {
	http.HandleFunc("/serve-id", idHandler)
	http.ListenAndServe(":8090", nil)
}
