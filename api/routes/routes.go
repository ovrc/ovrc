package routes

import (
	"github.com/go-chi/chi"
	"github.com/ovrc/ovrc/appcontext"
	"net/http"
)

// Resource holds the various context values.
type Resource struct {
	AppContext appcontext.AppContext
}

// SetRoutes sets all the routes for the API.
func (api Resource) SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", api.AuthLogin)
		r.Get("/logout", api.AuthLogout)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/me", api.UsersMe)
		r.Get("/", api.Users)
	})
	return r
}
