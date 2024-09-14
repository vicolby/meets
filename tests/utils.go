package tests

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
)

func setupTestContainer() (string, func(), error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", nil, err
	}

	host, err := postgresContainer.Host(context.Background())
	if err != nil {
		return "", nil, err
	}

	port, err := postgresContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		return "", nil, err
	}

	connString := fmt.Sprintf("host=%s port=%s user=testuser dbname=testdb password=testpass sslmode=disable", host, port.Port())

	terminate := func() {
		if err := postgresContainer.Terminate(context.Background()); err != nil {
			log.Fatalf("failed to terminate container: %v", err)
		}
	}

	return connString, terminate, nil
}
