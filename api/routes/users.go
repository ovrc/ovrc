package routes

import (
	"github.com/joaodlf/jsend"
	"net/http"
)

// UsersMe returns the logged in user details.
func (api Resource) UsersMe(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is valid.
	jsend.Write(w,
		jsend.StatusCode(200),
	)
}

func (api Resource) Users(w http.ResponseWriter, r *http.Request) {
	jsend.Write(w,
		jsend.StatusCode(200),
	)
}
