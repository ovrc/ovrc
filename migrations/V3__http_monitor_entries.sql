CREATE TABLE http_monitor_entries
(
    id                     SERIAL PRIMARY KEY,
    dt_created             TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    http_monitor_id        INTEGER                   NOT NULL,
    timeout                BOOLEAN     DEFAULT FALSE NOT NULL,
    dns_ms                 INTEGER,
    tls_handshake_ms       INTEGER,
    connect_ms             INTEGER,
    first_response_byte_ms INTEGER,
    total_ms               INTEGER
);