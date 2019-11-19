package models

import (
	"github.com/lib/pq"
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
