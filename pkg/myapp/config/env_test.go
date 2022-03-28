package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestRead(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		envVars map[string]string
		want    *Config
		wantErr bool
	}{
		{
			"test env vars",
			map[string]string{"MYAPP_SECRET": "supersecretkey"}, // pragma: allowlist secret
			&Config{Secret: "supersecretkey"},                   // pragma: allowlist secret
			false,
		},
		{
			"default env vars",
			map[string]string{},
			&Config{Secret: "12345"},
			false,
		},
	}
	for _, tt := range tests { // nolint:paralleltest
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				if err := os.Setenv(k, v); err != nil {
					t.Fatalf("failed to set env var: %v", err)
				}
			}
			got, err := Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			assert.Equal(t, got, tt.want)
			for k := range tt.envVars {
				if err := os.Unsetenv(k); err != nil {
					t.Fatalf("failed to set env var: %v", err)
				}
			}
		})
	}
}
