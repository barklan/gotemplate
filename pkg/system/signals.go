package system

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func HandleSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	log.Printf("received %s - exiting\n", sig)
	fmt.Println(sig)
	os.Exit(0)
}
