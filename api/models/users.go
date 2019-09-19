package models

import (
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"time"
)

// User represents teh users db table.
type User struct {
	Id        int         `db:"id"`
	DtCreated time.Time   `db:"dt_created"`
	DtUpdated pq.NullTime `db:"dt_updated"`
	Username  string      `db:"username"`
	Password  string      `db:"password"`
}

// SelectUser selects a single user via the username.
func (db *DB) SelectUser(username string) (*User, error) {
	user := &User{}

	err := db.Get(user, `SELECT * FROM users WHERE username = $1`, username)

	if err != nil {
		return user, errors.Wrap(err, "")
	}

	return user, nil
}
