package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/barklan/gotemplate/pkg/logging"
)

func handleSignals(sigs <-chan os.Signal) {
	sig := <-sigs
	log.Printf("received %s - exiting\n", sig)
	fmt.Println(sig)
	os.Exit(0)
}

func main() {
	lg := logging.Dev()
	defer func() {
		_ = lg.Sync()
	}()

	lg.Info("starting")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handleSignals(sigs)

	// Entry here
}
