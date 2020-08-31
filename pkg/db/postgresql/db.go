package postgresql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func (d *DB) Connect(dialect, dsn string) error {
	db, err := sqlx.Connect(dialect, dsn)
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}

func (d *DB) Migrate() error {
	tx, err := d.DB.Begin()

	_, err = tx.Exec(`create table if not exists schema_version (
		version text not null
		);

		create table if not exists users (
			id serial not null,
			username text not null unique,
			password text,
			is_admin bool default 'f',
			timezone text default 'UTC',
			last_login_at timestamp with time zone,
			primary key (id)
		);

		create table if not exists bookmarks (
			id bigserial not null,
			user_id int not null,
			title text not null,
			url text not null,
			created_at timestamp with time zone default now(),
			updated_at timestamp with time zone default now(),
			deleted_at timestamp with time zone default now(),
			primary key (id),
			unique (user_id, url),
			foreign key (user_id) references users(id) on delete cascade
		);

		create table if not exists tags (
			id bigserial not null,
			user_id int not null,
			name text not null,
			created_at timestamp with time zone default now(),
			updated_at timestamp with time zone default now(),
			deleted_at timestamp with time zone default now(),
			primary key (id),
			unique (user_id, name),
			foreign key (user_id) references users(id) on delete cascade
		);

		create table if not exists bookmark_tags (
			id bigserial not null,
			bookmark_id bigint not null,
			tag_id bigint not null,
			unique (bookmark_id, tag_id),
			foreign key (bookmark_id) references bookmarks(id) on delete cascade,
			foreign key (tag_id) references tags(id) on delete cascade
		);
		`)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (d *DB) Close() {
	d.DB.Close()
}
