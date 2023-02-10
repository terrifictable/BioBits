package main

import (
	"main/webserver"
	"github.com/ochinchina/go-ini"
	"fmt"
)

func main() {

	config := ini.Load("config.ini")
	token, err := config.GetValue("Config", "token")
	if err != nil {
		fmt.Println("[Warning]  No token in config file, you will not be able to use discord account embeds")
	}

	ws := webserver.NewWebServer(token)
	if err := ws.Start(2023); err != nil {
		panic(err)
	}
}
