package profile

import (
	"time"

	"go.uber.org/zap"
)

// To use: defer profile.Duration(time.Now(), "IntFactorial").
func Duration(lg *zap.Logger, invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	lg.Info("dirty profile", zap.String("name", name), zap.Duration("elapsed", elapsed))
}
