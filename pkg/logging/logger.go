package logging

import (
	"log"

	"github.com/barklan/gotemplate/pkg/system"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(iEnv system.InternalEnv) *zap.Logger {
	switch iEnv {
	case system.DevEnv:
		return Dev()
	case system.ProdEnv:
		return Prod()
	default:
		return Prod()
	}
}

func Dev() *zap.Logger {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.EncoderConfig.TimeKey = ""
	lg, err := zapConfig.Build()
	if err != nil {
		log.Fatal("failed to initialize logging")
	}

	return lg
}

func Prod() *zap.Logger {
	log.Println("prod logger not implemented, falling back to dev")
	return Dev()
}
