package dbtest

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// var DB *pgx.Conn // nolint:gochecknoglobals

func PrepareDB(migrationsPath string) (string, *dockertest.Pool, *dockertest.Resource) {
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
				"POSTGRES_DB=app",
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
	os.Setenv("POSTGRES_DB", "app")
	os.Setenv("POSTGRES_PASSWORD", "postgres")
	os.Setenv("POSTGRES_USER", "postgres")

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf(
		"postgres://postgres:postgres@%s/app?sslmode=disable",
		hostAndPort,
	)
	os.Setenv("POSTGRES_HOST_AND_PORT", hostAndPort)

	if err = resource.Expire(30); err != nil {
		log.Fatalf("failed to set expiration on container: %v", err)
	}

	log.Printf("connecting to database on %q\n", databaseUrl)
	pool.MaxWait = 180 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
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

	return databaseUrl, pool, resource
}
