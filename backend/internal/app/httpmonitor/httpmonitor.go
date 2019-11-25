package httpmonitor

import (
	"github.com/ovrc/ovrc/internal/appcontext"
	"log"
	"time"
)

// Resource holds the various context values.
type Resource struct {
	AppContext appcontext.AppContext
}

type requestDuration struct {
	start, connect, dns, tlsHandshake, total time.Duration
}

// Run kicks off the http monitors.
func (httpMon Resource) Run() {
	db := httpMon.AppContext.DB

	monitors, err := db.SelectHTTPMonitors()
	if err != nil {
		log.Fatalln("db:", err)
	}

	// TODO: All other http methods.
	for _, m := range monitors {
		switch m.Method {
		case "GET":
			go httpMon.get(m.Endpoint, m.ID)
		}
	}

	select {}
}
