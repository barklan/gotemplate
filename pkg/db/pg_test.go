package db

import (
	"log"
	"os"
	"testing"

	"github.com/barklan/gotemplate/pkg/dbtest"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	_, pool, resource := dbtest.PrepareDB("../../db/migrations")

	lg, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed to init zap logger")
	}

	_, err = Conn(lg)
	if err != nil {
		log.Fatalf("failed to connect to database")
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
