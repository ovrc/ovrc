package model

import (
	"time"
)

// HTTPMonitorEntry represents an entry in the http_monitor_entries db table.
type HTTPMonitorEntry struct {
	ID                  int       `db:"id"`
	DtCreated           time.Time `db:"dt_created"`
	HTTPMonitorID       int       `db:"http_monitor_id"`
	Timeout             bool      `db:"timeout"`
	DnsMs               int64     `db:"dns_ms"`
	TLSHandshakeMs      int64     `db:"tls_handshake_ms"`
	ConnectMs           int64     `db:"connect_ms"`
	FirstResponseByteMs int64     `db:"first_response_byte_ms"`
	TotalMs             int64     `db:"total_ms"`
}

func (db *DB) InsertHTTPMonitorEntry(entry HTTPMonitorEntry) (HTTPMonitorEntry, error) {
	_, err := db.NamedExec(`INSERT INTO http_monitor_entries (
						http_monitor_id, timeout, dns_ms, tls_handshake_ms, connect_ms, 
						first_response_byte_ms, total_ms) 
                        VALUES (:http_monitor_id, :timeout, :dns_ms, :tls_handshake_ms, 
                        :connect_ms, :first_response_byte_ms, :total_ms)`,
		&entry)

	if err != nil {
		return entry, err
	}

	return entry, nil
}
