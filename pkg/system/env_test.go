package system

import (
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGetInternalEnv(t *testing.T) {
	tests := []struct {
		name      string
		env_value string
		want      InternalEnv
		want1     bool
	}{
		{"dev environment", "dev", DevEnv, false},
		{"prod environment", "prod", ProdEnv, false},
		{"default environment", "", ProdEnv, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(InternalEnvKey, tt.env_value)
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
