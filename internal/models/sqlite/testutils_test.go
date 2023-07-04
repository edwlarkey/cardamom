package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	migratesqlite "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func newTestDB(t *testing.T, testName string) *sqlx.DB {
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?mode=memory&_time_format=sqlite", testName))
	if err != nil {
		t.Fatal(err)
	}
	dbx := sqlx.NewDb(db, "sqlite")

	migrations, err := iofs.New(MigrationFS, "migrations")
	if err != nil {
		t.Fatal(err)
	}
	driver, err := migratesqlite.WithInstance(db, &migratesqlite.Config{})
	if err != nil {
		t.Fatal(err)
	}
	m, err := migrate.NewWithInstance("iofs", migrations, "test", driver)
	if err != nil {
		t.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		dbx.Close()
	})

	return dbx
}
