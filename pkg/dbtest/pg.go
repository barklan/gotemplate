package dbtest

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ory/dockertest/v3" //nolint
	"github.com/ory/dockertest/v3/docker"

	// migrations.
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PrepareDB(migrationsPath string) (string, *dockertest.Pool, *dockertest.Resource) { //nolint:funlen
	var db *sql.DB
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Printf("starting database container\n")
	resource, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "14",
			Env: []string{
				"POSTGRES_PASSWORD=postgres",
				"POSTGRES_USER=postgres",
				"POSTGRES_DB=postgres",
				"listen_addresses = '*'",
			},
		},
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	os.Setenv("POSTGRES_DB", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "postgres")
	os.Setenv("POSTGRES_USER", "postgres")

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseURL := fmt.Sprintf(
		"postgres://postgres:postgres@%s/postgres?sslmode=disable",
		hostAndPort,
	)

	host := resource.GetBoundIP("5432/tcp")
	port := resource.GetPort("5432/tcp")
	os.Setenv("POSTGRES_HOST", host)
	os.Setenv("POSTGRES_PORT", port)

	if err = resource.Expire(30); err != nil {
		log.Fatalf("failed to set expiration on container: %v", err)
	}

	log.Printf("connecting to database on %q\n", databaseURL)
	pool.MaxWait = 180 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseURL)
		if err != nil {
			return fmt.Errorf("failed to open database connection: %w", err)
		}
		if err = db.Ping(); err != nil {
			return fmt.Errorf("failed to ping database: %w", err)
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	migrationManager, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"postgres",
		driver,
	)
	if err = migrationManager.Up(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return databaseURL, pool, resource
}
