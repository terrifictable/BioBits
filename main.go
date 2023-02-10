package main

import (
	"main/webserver"
)

func main() {

	ws := webserver.NewWebServer("OTE2MDkxMjU4MDEzODkyNjk5.GC2FP8.nMG5t7l2WnUBmSNNQBEhQgDEsYEbF-ANHRGq7U")
	if err := ws.Start(2023); err != nil {
		panic(err)
	}
}
