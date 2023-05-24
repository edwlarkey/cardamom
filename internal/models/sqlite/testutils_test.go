package sqlite

import (
	"database/sql"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func newTestDB(t *testing.T) *sqlx.DB {
	db, err := sql.Open("sqlite", "test.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	dbx := sqlx.NewDb(db, "sqlite")

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
		// os.Remove("test.sqlite")
	})

	return dbx
}
