package main

import (
	"log"

	"github.com/barklan/gotemplate/pkg/logging"
	"github.com/barklan/gotemplate/pkg/system"
	_ "go.uber.org/automaxprocs"
)

func main() {
	log.Println("starting myapp")
	go system.HandleSignals()

	logger := logging.NewAuto()
	// conf, err := config.Read()
	// if err != nil {
	// 	log.Panicf("failed to read config: %v\n", err)
	// }

	// Start app here
	logger.Info("main exited")
}
