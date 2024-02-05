package main

import (
	"flag"
	"fmt"
)

var (
	host  = flag.String("host", "localhost", "Host to listen on")
	port  = flag.String("port", "8080", "Port to listen on")
	proxy = flag.String("proxy", "http://localhost:4000", "Proxy to use")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%s", *host, *port)
	app := NewApp(AppSettings{
		Proxy: *proxy})
	err := app.Start(addr)
	if err != nil {
		fmt.Println(err)
	}
}
