package slog_test

import (
	"testing"

	"github.com/barklan/y/slog"
	"github.com/stretchr/testify/require"
)

func TestDev(t *testing.T) {
	_, err := slog.Dev()
	require.NoError(t, err)
}

func TestProd(t *testing.T) {
	_, err := slog.Prod()
	require.NoError(t, err)
}
