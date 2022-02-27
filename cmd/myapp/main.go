package main

import (
	"log"

	"github.com/barklan/gotemplate/pkg/system"
	_ "go.uber.org/automaxprocs"
)

func main() {
	log.Println("starting myapp")
	go system.HandleSignals()
	app, err := InitApp()
	if err != nil {
		log.Fatalf("app initialization failed: %v", err)
	}
	if err := app.Serve(); err != nil {
		log.Panic(err)
	}
}
