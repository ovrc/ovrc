package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"time"
)

// User represents teh users db table.
type User struct {
	ID        int         `db:"id"`
	DtCreated time.Time   `db:"dt_created"`
	DtUpdated pq.NullTime `db:"dt_updated"`
	Username  string      `db:"username"`
	Password  string      `db:"password"`
	SessionID uuid.UUID   `db:"session_id"`
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

func (db *DB) UpdateUserSessionID(userID int, sessionID uuid.UUID) error {
	_, err := db.NamedExec(`UPDATE users SET session_id=:session_id WHERE id=:id`,
		map[string]interface{}{
			"session_id": sessionID,
			"id":         userID,
		})

	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func (db *DB) SelectUserBySessionID(sessionID string) (*User, error) {
	user := &User{}

	err := db.Get(user, `SELECT * FROM users WHERE session_id = $1`, sessionID)

	if err != nil {
		return user, errors.Wrap(err, "")
	}

	return user, nil
}

func (db *DB) SelectUsersForAdmin() ([]User, error) {
	var users []User

	err := db.Select(&users, `SELECT username, dt_created FROM users`)

	if err != nil {
		return users, errors.Wrap(err, "")
	}

	return users, nil
}
