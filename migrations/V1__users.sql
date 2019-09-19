CREATE TABLE users(
    id serial PRIMARY KEY,
    dt_created TIMESTAMP DEFAULT NOW() NOT NULL,
    dt_updated TIMESTAMP,
    username text UNIQUE,
    password text
);