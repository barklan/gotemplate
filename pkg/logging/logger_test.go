// Package logging constructs zap loggers for different environments.
package logging

import (
	"os"
	"testing"
)

func TestNewAuto(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		envVars map[string]string
	}{
		{
			"logger for dev environment",
			map[string]string{"INTERNAL_ENVIRONMENT": "dev"},
		},
		{
			"logger for prod environment",
			map[string]string{"INTERNAL_ENVIRONMENT": ""},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}
			if _, err := NewAuto(); err != nil {
				t.Fatalf("error when constructing logger: %v\n", err)
			}
		})
	}
}
