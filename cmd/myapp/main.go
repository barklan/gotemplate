package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func handleSignals(sigs <-chan os.Signal, done chan<- struct{}) {
	sig := <-sigs
	log.Printf("received %s - exiting\n", sig)
	fmt.Println(sig)
	os.Exit(0)
}

func main() {
	log.Println("starting...")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{}, 1)
	go handleSignals(sigs, done)

	// Entry here
}
