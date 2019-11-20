package appcontext

import "github.com/ovrc/ovrc/internal/model"

type AppContext struct {
	DB     model.Datastore
	Config ConfigSpecification
}

type ConfigSpecification struct {
	Env               string
	UseSSL            string   `envconfig:"USE_SSL"`
	WebPort           string   `envconfig:"WEB_PORT"`
	WebCertFile       string   `envconfig:"WEB_CERT_FILE"`
	WebKeyFile        string   `envconfig:"WEB_KEY_FILE"`
	WebAllowedOrigins []string `envconfig:"WEB_ALLOWED_ORIGINS"`
	DBConnection      string   `envconfig:"DB_CONNECTION"`
}
