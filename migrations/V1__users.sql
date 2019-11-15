CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    dt_created TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    dt_updated TIMESTAMPTZ,
    username   TEXT UNIQUE,
    password   TEXT,
    session_id UUID UNIQUE
);