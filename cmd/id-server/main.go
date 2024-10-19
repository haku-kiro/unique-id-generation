package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	gen "unique.ids/internal/id-gen"
)

func main() {
	incrementer := new(gen.Inc)
	rpc.Register(incrementer)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http.Serve(l, nil)
}
