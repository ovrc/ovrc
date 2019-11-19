CREATE TABLE http_monitors
(
    id         SERIAL PRIMARY KEY,
    dt_created TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    dt_updated TIMESTAMPTZ,
    endpoint   TEXT,
    method     TEXT
);