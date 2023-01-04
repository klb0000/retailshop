package main

import (
	"fmt"

	"github.com/klb0000/retailshop/server"
)

func main() {
	addr := "localhost:8080"
	fmt.Printf("starting server at %s\n", addr)
	server.Serve(addr)

}
