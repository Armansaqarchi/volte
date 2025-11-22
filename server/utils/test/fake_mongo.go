package test

import (
	"fmt"
	"log/slog"
	"testing"
	"volte/backend/databases"

	"github.com/docker/docker/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	testUsername = "test"
	testPassword = "test"
	testHost     = "localhost"
	testPort     = "27019"
)

var mongoServer *testcontainers.Container

func CreateNewFakeMongoServer(t *testing.T) *testcontainers.Container {
	slog.Info("Creating container for mongo server.")
	if mongoServer != nil {
		return mongoServer
	}
	slog.Info("Initiating a container request for mongo.")
	mongoContainerRequest := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{fmt.Sprintf("%s:%s/tcp", testPort, "27017")},
		ConfigModifier: func(config *container.Config) {
			config.Hostname = testHost
		},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": testUsername,
			"MONGO_INITDB_ROOT_PASSWORD": testPassword,
		},
		WaitingFor: wait.ForLog("Waiting for connections"),
	}
	slog.Info("Creating mongo container.")
	mongoContainer, err := testcontainers.GenericContainer(t.Context(), testcontainers.GenericContainerRequest{
		ContainerRequest: mongoContainerRequest,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	mongoServer = &mongoContainer
	slog.Info("Successfully created mongo container.")
	return mongoServer
}

func NewFakeMongoClient() *databases.MongoClient {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", testUsername, testPassword, testHost, testPort)
	slog.Info(fmt.Sprintf("Connecting to fake mongo server at %s", url))
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	return databases.NewMongoClientWithConfig(opts)
}
