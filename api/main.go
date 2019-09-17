package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joaodlf/jsend"
	"github.com/spf13/viper"
	"github.com/teamwork/reload"
	"log"
	"net/http"
)

func main() {
	// Load config from file.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Auto reload on build.
	if viper.GetString("env") == "development" {
		go func() {
			err := reload.Do(log.Printf)
			if err != nil {
				panic(err) // Only returns initialisation errors.
			}
		}()
	}

	// Web server.
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS middleware so that the frontend can communicate with this API.
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("webserver.allowed_origins"),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)

	// Register routes.
	r.Mount("/auth", authRouter())
	r.Mount("/users", usersRouter())

	// Serve over HTTPS.
	http.ListenAndServeTLS(viper.GetString("webserver.port"),
		viper.GetString("webserver.cert_file"),
		viper.GetString("webserver.key_file"),
		r)
}

func authRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/login", authLogin)
	return r
}

func usersRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/me", usersMe)
	return r
}

type AuthLoginForm struct {
	Username string
	Password string
	Errors   map[string]interface{}
}

func (form *AuthLoginForm) Validate() bool {
	form.Errors = make(map[string]interface{})

	if form.Username == "" {
		form.Errors["username"] = "missing"
	}

	if form.Password == "" {
		form.Errors["password"] = "missing"
	}

	return len(form.Errors) == 0
}

func authLogin(w http.ResponseWriter, r *http.Request) {
	form := &AuthLoginForm{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	if form.Validate() == false {
		jsend.Write(w,
			jsend.Data(form.Errors),
			jsend.StatusCode(400),
		)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "test",
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	jsend.Write(w)
}

func usersMe(w http.ResponseWriter, r *http.Request) {
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
