CREATE TABLE records (
  id serial PRIMARY KEY,
  bin_id varchar(50) NOT NULL,
  hook_id varchar(50) NOT NULL,
  date timestamp NOT NULL
);