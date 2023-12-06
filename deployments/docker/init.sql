CREATE USER pguser WITH PASSWORD 'pgpwd';
GRANT ALL PRIVILEGES ON DATABASE "testdb" TO pguser;

CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'postgres';

CREATE TABLE IF NOT EXISTS "projects" (
  id integer PRIMARY KEY,
  name TEXT,
  description TEXT,
  avatar_url TEXT,
  type TEXT,
  archived BOOLEAN
);

CREATE TABLE IF NOT EXISTS "authors" (
  id serial PRIMARY KEY,
  name TEXT,
  display_name TEXT
);

CREATE TABLE IF NOT EXISTS "issues" (
  id integer PRIMARY KEY,
  project_id integer REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE,
  author_id integer REFERENCES authors (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  reporter_id integer REFERENCES authors (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  key TEXT,
  summary TEXT,
  type TEXT,
  priority TEXT,
  closed BOOLEAN,
  created_time TIMESTAMP WITHOUT TIME ZONE,
  closed_time TIMESTAMP WITHOUT TIME ZONE,
  updated_time TIMESTAMP WITHOUT TIME ZONE,
  time_spent INT
);
