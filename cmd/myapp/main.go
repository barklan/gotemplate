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

	logger, err := logging.NewAuto()
	if err != nil {
		log.Fatalf("failed to init logger: %v\n", err)
	}
	// conf, err := config.Read()
	// if err != nil {
	// 	log.Panicf("failed to read config: %v\n", err)
	// }

	// Start app here
	logger.Info("main exited")
}
