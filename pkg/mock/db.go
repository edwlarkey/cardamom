package mock

type DB struct{}

func (d *DB) Connect(dialect string, dsn string) error {
	return nil
}

func (d *DB) Migrate() error {
	return nil
}

func (d *DB) Close() {
}
