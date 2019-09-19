package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/ovrc/ovrc/appcontext"
	"github.com/ovrc/ovrc/models"
	"github.com/ovrc/ovrc/routes"
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

	db, err := models.NewDB("user=ovrc dbname=ovrc password=ovrc sslmode=disable")
	if err != nil {
		log.Fatalln(err)
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

	ac := appcontext.AppContext{DB: db}

	// Register routes.
	api := routes.Resource{AppContext: ac}
	r.Mount("/", api.SetRoutes())

	// Serve over HTTPS.
	http.ListenAndServeTLS(viper.GetString("webserver.port"),
		viper.GetString("webserver.cert_file"),
		viper.GetString("webserver.key_file"),
		r)
}
