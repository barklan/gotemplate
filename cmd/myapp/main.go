package main

import (
	"github.com/barklan/gotemplate/pkg/logging"
	"github.com/barklan/gotemplate/pkg/myapp"
	"github.com/barklan/gotemplate/pkg/system"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {
	go system.HandleSignals()
	internalEnv, _ := system.GetInternalEnv()

	lg := logging.New(internalEnv)
	defer func() {
		_ = lg.Sync()
	}()
	lg.Info("starting")
	defer lg.Info("exiting now")

	if err := myapp.Start(lg); err != nil {
		lg.Panic("error in myapp", zap.Error(err))
	}
}
