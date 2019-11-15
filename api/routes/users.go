package routes

import (
	"github.com/joaodlf/jsend"
	"github.com/ovrc/ovrc/models"
	"net/http"
)

// UsersMe returns the logged in user details.
func (api Resource) UsersMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)

	jsend.Write(w,
		jsend.StatusCode(200),
		jsend.Data(map[string]interface{}{
			"username": user.Username,
		}),
	)
}

func (api Resource) Users(w http.ResponseWriter, r *http.Request) {
	jsend.Write(w,
		jsend.StatusCode(200),
	)
}
