package caching

import (
	"context"
	"time"
)

var (
	ctx               = context.TODO()
	variableKeySymbol = "."
)

type Cache interface {
	Set(string, interface{}, time.Duration) error
	Get(string) ([]byte, bool, error)
	SetVar(string, string, interface{}, time.Duration) error
	GetVar(string, string) ([]byte, bool, error)
}

type FastCache interface {
	Set(interface{}, interface{})
	Get(interface{}) (interface{}, bool)
}
