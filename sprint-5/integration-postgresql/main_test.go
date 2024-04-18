package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	var db *pgx.Conn

	ctx := context.Background()
	dbName := "postgres"
	dbUser := "postgres"
	dbPassword := "postgres"

	pcDB, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:14-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("could not start container: %v", err)
	}

	// Get container host and port
	host, err := pcDB.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get PostgreSQL container host: %s", err)
	}
	port, err := pcDB.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Could not get PostgreSQL container port: %s", err)
	}

	connString := fmt.Sprintf("user=postgres password=postgres dbname=postgres host=%s port=%s sslmode=disable", host, port.Port())
	db, err = pgx.Connect(ctx, connString)
	if err != nil {
		return
	}

	code := m.Run()
	err = db.Close(ctx)
	if err != nil {
		log.Panicf("Could not close PostgreSQL container: %s", err)
	}
	os.Exit(code)
}

func TestQueryExecution(t *testing.T) {
	var db *pgx.Conn
	if &db == nil { //nolint:govet,staticcheck // it's not my code
		require.FailNow(t, "db is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert test data
	_, err := db.Exec(ctx, "CREATE TABLE test_table (id SERIAL PRIMARY KEY, name VARCHAR(255))")
	require.NoError(t, err)
	_, err = db.Exec(ctx, "INSERT INTO test_table (name) VALUES ('test_name')")
	require.NoError(t, err)

	// Perform query
	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM test_table").Scan(&count)
	require.NoError(t, err)

	// Assert the result
	assert.Equal(t, 1, count)
}
