//go:build wireinject
// +build wireinject

package main

import (
	"github.com/barklan/gotemplate/pkg/caching"
	"github.com/barklan/gotemplate/pkg/db"
	"github.com/barklan/gotemplate/pkg/logging"
	"github.com/barklan/gotemplate/pkg/myapp/config"
	"github.com/barklan/gotemplate/pkg/myapp/rest"
	"github.com/google/wire"
)

func InitApp() (*rest.PublicCtrl, error) {
	wire.Build(logging.NewAuto, config.Read, db.Conn, caching.NewArc, rest.NewCtrl)
	return &rest.PublicCtrl{}, nil
}
