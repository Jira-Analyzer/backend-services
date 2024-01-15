CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'postgres';

CREATE TABLE IF NOT EXISTS "Project" (
  id integer PRIMARY KEY,
  name TEXT,
  description TEXT,
  avatar_url TEXT,
  type TEXT,
  archived BOOLEAN
);

CREATE TABLE IF NOT EXISTS "Author" (
  id serial PRIMARY KEY,
  name TEXT,
  display_name TEXT
);

CREATE TABLE IF NOT EXISTS "Issue" (
  id integer PRIMARY KEY,
  project_id integer REFERENCES "Project" (id) ON DELETE CASCADE ON UPDATE CASCADE,
  author_id integer REFERENCES "Author" (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  reporter_id integer REFERENCES "Author" (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  key TEXT,
  summary TEXT,
  type TEXT,
  priority TEXT,
  status TEXT,
  created_time TIMESTAMP WITHOUT TIME ZONE,
  closed_time TIMESTAMP WITHOUT TIME ZONE,
  updated_time TIMESTAMP WITHOUT TIME ZONE
);
