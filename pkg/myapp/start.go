package myapp

import (
	"fmt"

	"github.com/barklan/gotemplate/pkg/caching"
	"github.com/barklan/gotemplate/pkg/db"
	"github.com/barklan/gotemplate/pkg/myapp/config"
	"github.com/barklan/gotemplate/pkg/myapp/rest"
	"go.uber.org/zap"
)

func Start(lg *zap.Logger) error {
	conf, err := config.Read()
	if err != nil {
		return fmt.Errorf("failed to read config for myapp: %w", err)
	}

	db, err := db.Conn(lg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	fc, err := caching.NewArc(128)
	if err != nil {
		return fmt.Errorf("failed to init arc cache: %w", err)
	}

	ctrl := rest.NewCtrl(lg, conf, db, fc)
	return ctrl.Serve()
}
