package postgres_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/edwlarkey/cardamom/pkg/db/postgres"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var (
	user     = "cardamom"
	password = "secret"
	database = "cardamom"
	port     = "5433"
	dialect  = "postgres"
	dsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
	idleConn = 25
	maxConn  = 25
)

var db = &postgres.DB{}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.3",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + database,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	dsn = fmt.Sprintf(dsn, user, password, port, database)
	if err = pool.Retry(func() error {
		err = db.Connect(dialect, dsn)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	defer func() {
		db.Close()
	}()

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestMigrate(t *testing.T) {
	err := db.Migrate()
	if err != nil {
		t.Errorf("Inserting User failed %s", err)
	}
}

func TestMigrateAgain(t *testing.T) {
	err := db.Migrate()
	if err != nil {
		t.Errorf("Inserting User failed %s", err)
	}
}
