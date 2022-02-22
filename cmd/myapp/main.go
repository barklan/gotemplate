package main

import (
	"log"

	"github.com/barklan/gotemplate/pkg/system"
	_ "go.uber.org/automaxprocs"
)

// func main() {
// 	go system.HandleSignals()
// 	internalEnv, _ := system.GetInternalEnv()

// 	lg := logging.New(internalEnv)
// 	defer func() {
// 		_ = lg.Sync()
// 	}()
// 	lg.Info("starting")

// 	if err := myapp.Start(lg); err != nil {
// 		lg.Panic("error in myapp", zap.Error(err))
// 	}
// }

func main() {
	go system.HandleSignals()
	app, err := InitApp()
	if err != nil {
		log.Fatalf("app initialization failed: %v", err)
	}
	app.Serve()
}
