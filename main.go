package main

import "go-starter/domain/app"

func main() {
	app, err := app.New()
	if err != nil {
		panic(err)
	}
	println("Server listening on " + app.Server.Addr)
	app.Server.ListenAndServe()
}
