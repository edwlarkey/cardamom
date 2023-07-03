package sqlite

import (
	"database/sql"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func newTestDB(t *testing.T) *sqlx.DB {
	db, err := sql.Open("sqlite", "test.sqlite?_time_format=sqlite")
	if err != nil {
		t.Fatal(err)
	}
	dbx := sqlx.NewDb(db, "sqlite")

	migrations, err := iofs.New(MigrationFS, "migrations")
	if err != nil {
		t.Fatal(err)
	}
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
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
		err = m.Down()
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
		os.Remove("test.sqlite")
	})

	return dbx
}
