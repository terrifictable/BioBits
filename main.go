package main

import (
	"main/webserver"
)

func main() {

	ws := webserver.NewWebServer("")
	if err := ws.Start(2023); err != nil {
		panic(err)
	}
}
