package rest

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/barklan/gotemplate/pkg/caching"
	"github.com/barklan/gotemplate/pkg/db"
	"github.com/barklan/gotemplate/pkg/dbtest"
	"github.com/barklan/gotemplate/pkg/logging"
	"github.com/barklan/gotemplate/pkg/myapp/config"
	"github.com/joho/godotenv"
)

func MockCtrl() (*PublicCtrl, error) {
	conf, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to construct test config: %w", err)
	}
	lg := logging.Dev()
	db, err := db.Conn(lg)
	if err != nil {
		return nil, fmt.Errorf("failed to open test db conn: %w", err)
	}
	fc, err := caching.NewArc(conf)
	if err != nil {
		return nil, fmt.Errorf("faield to init arc cache: %w", err)
	}

	ctrl := NewCtrl(lg, conf, db, fc)

	return ctrl, nil
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	if testing.Short() {
		os.Exit(m.Run())
	} else {
		_, pool, resource := dbtest.PrepareDB("../../../db/migrations")
		code := m.Run()
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
		os.Exit(code)
	}
}
