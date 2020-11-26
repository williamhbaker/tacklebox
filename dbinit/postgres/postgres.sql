CREATE TABLE users (
  id serial PRIMARY KEY,
  email varchar(255) UNIQUE NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created timestamp NOT NULL
);

CREATE TABLE bins (
  id varchar(50) PRIMARY KEY,
  user_id integer REFERENCES users (id),
  created timestamp NOT NULL
);

CREATE TABLE records (
  id serial PRIMARY KEY,
  bin_id varchar(50) REFERENCES bins (id),
  hook_id varchar(50) NOT NULL,
  created timestamp NOT NULL
);