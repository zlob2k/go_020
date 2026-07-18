package main

// cmd

import (
	"log"

	app "example.com/zlob2k/go_020/internal/app"
)

func main() {
	context := app.NewContext()
	err := context.WebServer().Start()
	if err != nil {
		log.Fatalf("\nError starting server %s \n%v", context.ServerAddr, err)
	}
	defer context.Shutdown()
}
