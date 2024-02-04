package main

import (
	"flag"
	"fmt"
	"go-starter/router"
	"net/http"
)

var (
	host = flag.String("host", "localhost", "Host to listen on")
	port = flag.Int("port", 8080, "Port to listen on")
)

func main() {
	flag.Parse()
	router := router.New()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Listening on %s...\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Println(err)
	}
}
