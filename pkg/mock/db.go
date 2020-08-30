package mock

type DB struct{}

func (d *DB) Connect(dbtype string, dsn string) error {
	return nil
}

func (d *DB) Migrate() error {
	return nil
}

func (d *DB) Close() {
}
