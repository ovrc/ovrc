package model

import (
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"time"
)

// HTTPMonitor represents the http_monitors db table.
type HTTPMonitor struct {
	ID        int         `db:"id"`
	DtCreated time.Time   `db:"dt_created"`
	DtUpdated pq.NullTime `db:"dt_updated"`
	Endpoint  string      `db:"endpoint"`
	Method    string      `db:"method"`
}

// SelectHTTPMonitors selects all HTTP monitors.
func (db *DB) SelectHTTPMonitors() ([]HTTPMonitor, error) {
	var monitors []HTTPMonitor

	err := db.Select(&monitors, `SELECT * FROM http_monitors`)

	if err != nil {
		return monitors, errors.Wrap(err, "")
	}

	return monitors, nil
}
