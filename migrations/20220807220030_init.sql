-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id  INTEGER NOT NULL,
  name  TEXT NOT NULL,
  email TEXT NOT NULL,
  hashed_password TEXT NOT NULL,
  created DATETIME NOT NULL,
  PRIMARY KEY(id AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS bookmarks (
  id  INTEGER NOT NULL,
  title  TEXT NOT NULL,
  url TEXT NOT NULL,
  content TEXT,
  user_id INTEGER NOT NULL,
  created DATETIME NOT NULL,
  PRIMARY KEY(id AUTOINCREMENT)
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE sessions (
  token TEXT PRIMARY KEY,
  data BLOB NOT NULL,
  expiry REAL NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);

CREATE UNIQUE INDEX "uni_user_email" ON "bookmarks" (
  "email"
);

CREATE UNIQUE INDEX "uni_user_title" ON "bookmarks" (
  "title",
  "user_id"
);

CREATE UNIQUE INDEX "uni_user_url" ON "bookmarks" (
  "url",
  "user_id"
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE 'sessions';
DROP TABLE 'bookmarks';
DROP TABLE 'users';
-- +goose StatementEnd
