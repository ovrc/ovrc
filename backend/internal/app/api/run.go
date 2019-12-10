package api

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joaodlf/jsend"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/teamwork/reload"
	"log"
	"net/http"
)

// SessionCheck is a middleware to check for the session_id before allowing access to the API.
func SessionCheck(db *model.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Allow /auth/login through so that the user can actually login.
			if r.URL.String() == "/auth/login" {
				next.ServeHTTP(w, r)
				return
			}

			sessionID, err := r.Cookie("session_id")

			if err != nil {
				jsend.Write(w,
					jsend.StatusCode(403),
					jsend.Data(map[string]interface{}{
						"validation": "Missing session_id.",
					}),
				)
				return
			}

			user, err := db.SelectUserBySessionID(sessionID.Value)

			if err != nil {
				jsend.Write(w,
					jsend.StatusCode(403),
					jsend.Data(map[string]interface{}{
						"validation": "Could not validate session_id.",
					}),
				)
				return
			}

			// Put the user in the current request context.
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Run() {
	var config appcontext.ConfigSpecification

	err := envconfig.Process("ovrc", &config)
	if err != nil {
		log.Fatal("config:", err.Error())
	}

	// Auto reload on build.
	if config.Env == "development" {
		go func() {
			err := reload.Do(log.Printf)
			if err != nil {
				panic(err) // Only returns initialisation errors.
			}
		}()
	}
	db, err := model.NewDB(config.DBConnection)
	if err != nil {
		log.Fatalln("db:", err)
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
		AllowedOrigins:   config.WebAllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)
	r.Use(SessionCheck(db))

	ac := appcontext.AppContext{DB: db, Config: config}

	// Register routes.
	apiResource := Resource{AppContext: ac}
	r.Mount("/", apiResource.SetRoutes())

	// Not currently in use - but something to grow from in the future.
	if config.UseSSL == "true" {
		err = http.ListenAndServeTLS(config.WebPort,
			config.WebCertFile,
			config.WebKeyFile,
			r)
	} else {
		err = http.ListenAndServe(config.WebPort, r)
	}

	if err != nil {
		log.Fatalln("web:", err)
	}
}
