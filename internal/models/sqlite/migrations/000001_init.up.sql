CREATE TABLE IF NOT EXISTS users (
  id  INTEGER NOT NULL,
  name  TEXT NOT NULL,
  email TEXT NOT NULL,
  hashed_password TEXT NOT NULL,
  created DATETIME NOT NULL,
  PRIMARY KEY(id AUTOINCREMENT)
);

CREATE UNIQUE INDEX "uni_user_email" ON "users" (
  "email"
);

CREATE TABLE IF NOT EXISTS bookmarks (
  id  INTEGER NOT NULL,
  title  TEXT NOT NULL,
  url TEXT NOT NULL,
  content TEXT,
  read INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  created DATETIME NOT NULL,
  PRIMARY KEY(id AUTOINCREMENT)
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX "uni_user_title" ON "bookmarks" (
  "title",
  "user_id"
);

CREATE UNIQUE INDEX "uni_user_url" ON "bookmarks" (
  "url",
  "user_id"
);

CREATE TABLE IF NOT EXISTS tags (
  id  INTEGER NOT NULL,
  name  TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  created DATETIME NOT NULL,
  PRIMARY KEY(id AUTOINCREMENT)
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX "uni_user_tag" ON "tags" (
  "name",
  "user_id"
);

CREATE TABLE IF NOT EXISTS bookmark_tags (
  id  INTEGER NOT NULL,
  bookmark_id INTEGER NOT NULL,
  tag_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  PRIMARY KEY(id AUTOINCREMENT)
  FOREIGN KEY(bookmark_id) REFERENCES bookmarks(id)
  FOREIGN KEY(tag_id) REFERENCES tags(id)
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX "uni_bookmark_tag" ON "bookmark_tags" (
  "bookmark_id",
  "tag_id",
  "user_id"
);

CREATE TABLE sessions (
  token TEXT PRIMARY KEY,
  data BLOB NOT NULL,
  expiry REAL NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);
