package main

import (
	"github.com/barklan/gotemplate/pkg/logging"
	"github.com/barklan/gotemplate/pkg/system"
)

func main() {
	go system.HandleSignals()
	internalEnv, _ := system.GetInternalEnv()

	lg := logging.New(internalEnv)
	defer func() {
		_ = lg.Sync()
	}()
	lg.Info("starting")
	defer lg.Warn("main exited")

	// Entry here
}
