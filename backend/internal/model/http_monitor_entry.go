package model

import (
	"github.com/pkg/errors"
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

// HTTPMonitorEntryDashboard represents a row for the HTTP Monitoring Dashboard.
type HTTPMonitorEntryDashboard struct {
	ID         int    `db:"id"`
	Endpoint   string `db:"endpoint"`
	Method     string `db:"method"`
	AvgTotalMs int    `db:"avg_total_ms"`
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

func (db *DB) SelectHTTPMonitorEntriesForDashboard() ([]HTTPMonitorEntryDashboard, error) {
	var entries []HTTPMonitorEntryDashboard

	err := db.Select(&entries, `SELECT hm.id, hm.endpoint, hm.method, cast(avg(hme.total_ms) AS INT) AS avg_total_ms
							FROM http_monitor_entries AS hme
         					JOIN http_monitors AS hm ON hm.id = hme.http_monitor_id
							WHERE hme.dt_created >= current_timestamp - INTERVAL '5 hours'
							GROUP BY hm.id, hm.endpoint, hm.method
							ORDER BY hm.id ASC;`)

	if err != nil {
		return entries, errors.Wrap(err, "")
	}

	return entries, nil
}
