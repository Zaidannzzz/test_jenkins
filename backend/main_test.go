package main_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	// Start Docker container
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("postgres", "latest", []string{"POSTGRES_PASSWORD=Samsunggalaxyy1"})
	if err != nil {
		log.Fatalf("Could not start Docker container: %s", err)
	}

	// Get container IP and port
	host := resource.GetHostPort("5432/tcp")

	// Wait for the container to be ready
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=postgres password=Samsunggalaxyy1 dbname=testpostgre sslmode=disable", host))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to Docker container: %s", err)
	}

	// Run tests
	exitCode := m.Run()

	// Clean up Docker container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge Docker container: %s", err)
	}

	os.Exit(exitCode)
}

func TestDatabaseConnection(t *testing.T) {
	assert.NotNil(t, db, "Database connection is not established")
}
