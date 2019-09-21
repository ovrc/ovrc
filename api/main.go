package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/ovrc/ovrc/appcontext"
	"github.com/ovrc/ovrc/models"
	"github.com/ovrc/ovrc/routes"
	"github.com/teamwork/reload"
	"log"
	"net/http"
)

func main() {
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

	db, err := models.NewDB(config.DBConnection)
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

	ac := appcontext.AppContext{DB: db, Config: config}

	// Register routes.
	api := routes.Resource{AppContext: ac}
	r.Mount("/", api.SetRoutes())

	// Serve over HTTPS.
	err = http.ListenAndServeTLS(config.WebPort,
		config.WebCertFile,
		config.WebKeyFile,
		r)

	if err != nil {
		log.Fatalln("web:", err)
	}
}
