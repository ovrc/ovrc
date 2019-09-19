package routes

import (
	"github.com/joaodlf/jsend"
	"net/http"
)

// UsersMe returns the logged in user details.
func (api Resource) UsersMe(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_id")

	if err != nil {
		jsend.Write(w,
			jsend.StatusCode(400),
		)
		return
	}

	jsend.Write(w,
		jsend.StatusCode(200),
	)
}
