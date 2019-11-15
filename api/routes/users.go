package routes

import (
	"fmt"
	"github.com/joaodlf/jsend"
	"github.com/ovrc/ovrc/models"
	"net/http"
)

// UsersMe returns the logged in user details.
func (api Resource) UsersMe(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is valid.
	user := r.Context().Value("user").(*models.User)

	fmt.Println("Logged in user:", user)

	jsend.Write(w,
		jsend.StatusCode(200),
	)
}

func (api Resource) Users(w http.ResponseWriter, r *http.Request) {
	jsend.Write(w,
		jsend.StatusCode(200),
	)
}
