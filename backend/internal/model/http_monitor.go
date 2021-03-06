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

func (db *DB) InsertHTTPMonitor(mon HTTPMonitor) (HTTPMonitor, error) {
	res := HTTPMonitor{}
	stmt, err := db.PrepareNamed(`INSERT INTO http_monitors (method, endpoint) 
                        	VALUES (:method, :endpoint) RETURNING *`)

	if err != nil {
		return res, err
	}

	err = stmt.Get(&res, mon)

	if err != nil {
		return res, err
	}

	return res, nil
}

// CountActiveHTTPMonitors returns the total active HTTP Monitors.
func (db *DB) CountActiveHTTPMonitors() (int, error) {
	count := 0
	row := db.QueryRow(`SELECT count(*) FROM http_monitors`)

	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
