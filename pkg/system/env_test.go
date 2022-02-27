package system

import (
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) { //nolint:interfacer
	goleak.VerifyTestMain(m)
}

func TestGetInternalEnv(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		envValue string
		want     InternalEnv
		want1    bool
	}{
		{"dev environment", "dev", DevEnv, false},
		{"prod environment", "prod", ProdEnv, false},
		{"default environment", "", ProdEnv, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			os.Setenv(InternalEnvKey, tt.envValue)
			got, got1 := GetInternalEnv()
			if got != tt.want {
				t.Errorf("GetInternalEnv() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetInternalEnv() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
