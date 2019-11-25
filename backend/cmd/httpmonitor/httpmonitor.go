package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/ovrc/ovrc/internal/app/httpmonitor"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/teamwork/reload"
	"log"
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

	db, err := model.NewDB(config.DBConnection)
	if err != nil {
		log.Fatalln("db:", err)
	}

	ac := appcontext.AppContext{DB: db, Config: config}

	httpMonResource := httpmonitor.Resource{AppContext: ac}

	httpMonResource.Run()
}
