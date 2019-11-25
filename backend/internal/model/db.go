package model

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Datastore represents all the queries.
type Datastore interface {
	SelectUser(username string) (*User, error)
	UpdateUserSessionID(userID int, sessionID uuid.UUID) error
	SelectUserBySessionID(sessionID string) (*User, error)
	SelectUsersForAdmin() ([]User, error)

	SelectHTTPMonitors() ([]HTTPMonitor, error)

	InsertHTTPMonitorEntry(HTTPMonitorEntry) (HTTPMonitorEntry, error)
	SelectHTTPMonitorEntriesForDashboard() ([]HTTPMonitorEntryDashboard, error)
}

// DB holds a sql db.
type DB struct {
	*sqlx.DB
}

// NewDB opens the database and pings the server to test the connection.
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
