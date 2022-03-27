package config

import (
	"os"
	"reflect"
	"testing"

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
			"test env var",
			map[string]string{"MYAPP_SECRET": "supersecretkey"}, // pragma: allowlist secret
			&Config{Secret: "supersecretkey"},                   // pragma: allowlist secret
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
