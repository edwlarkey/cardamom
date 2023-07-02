package sqlite

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

//go:embed migrations
var migrationFS embed.FS

func OpenDB(dsn string) (*sqlx.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	dbx := sqlx.NewDb(db, "sqlite")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = migrateDB(db); err != nil {
		return nil, err
	}
	return dbx, nil
}

func migrateDB(db *sql.DB) error {
	migrations, err := iofs.New(migrationFS, "migrations")
	if err != nil {
		return err
	}
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", migrations, "cardamom", driver)
	if err != nil {
		return err
	}
	m.Up()

	return nil
}
