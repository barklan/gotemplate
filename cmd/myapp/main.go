package main

import (
	"log"

	"github.com/barklan/gotemplate/pkg/system"
	_ "go.uber.org/automaxprocs"
)

func main() {
	go system.HandleSignals()
	app, err := InitApp()
	if err != nil {
		log.Fatalf("app initialization failed: %v", err)
	}
	app.Serve()
}
