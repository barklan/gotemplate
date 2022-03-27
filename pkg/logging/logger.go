// Package logging constructs zap loggers for different environments.
package logging

import (
	"github.com/barklan/gotemplate/pkg/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewAuto() (*zap.Logger, error) {
	internalEnv, _ := system.GetInternalEnv()

	return New(internalEnv)
}

func New(iEnv system.InternalEnv) (*zap.Logger, error) {
	switch iEnv {
	case system.DevEnv:
		return Dev()
	case system.ProdEnv:
		return Prod()
	default:
		return Prod()
	}
}

func Dev() (*zap.Logger, error) {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.EncoderConfig.TimeKey = ""
	return zapConfig.Build()
}

func Prod() (*zap.Logger, error) {
	return zap.NewProduction()
}
