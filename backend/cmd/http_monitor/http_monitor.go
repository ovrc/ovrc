package main

import (
	"crypto/tls"
	"github.com/kelseyhightower/envconfig"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/teamwork/reload"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
	"time"
)

type requestDuration struct {
	start, connect, dns, tlsHandshake, total time.Duration
}

func timeGet(db model.DB, url string, mID int) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case _ = <-ticker.C:
			req, _ := http.NewRequest("GET", url, nil)

			var connect, dns, tlsHandshake, start time.Time
			rd := requestDuration{}

			trace := &httptrace.ClientTrace{
				DNSStart: func(dsi httptrace.DNSStartInfo) {
					dns = time.Now()
				},
				DNSDone: func(ddi httptrace.DNSDoneInfo) {
					rd.dns = time.Since(dns)
				},

				TLSHandshakeStart: func() {
					tlsHandshake = time.Now()
				},
				TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
					rd.tlsHandshake = time.Since(tlsHandshake)
				},

				ConnectStart: func(network, addr string) {
					connect = time.Now()
				},
				ConnectDone: func(network, addr string, err error) {
					rd.connect = time.Since(connect)
				},

				GotFirstResponseByte: func() {
					rd.start = time.Since(start)
				},
			}

			req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

			// Create a new http.Transport to avoid caching results.
			client := &http.Client{Transport: &http.Transport{}, Timeout: 10 * time.Second}

			start = time.Now()
			_, err := client.Do(req)
			rd.total = time.Since(start)

			entry := model.HTTPMonitorEntry{
				HTTPMonitorID: mID,
			}

			if err, ok := err.(net.Error); ok && err.Timeout() {
				entry.Timeout = true
			}
			entry.DnsMs = rd.dns.Milliseconds()
			entry.TLSHandshakeMs = rd.tlsHandshake.Milliseconds()
			entry.ConnectMs = rd.connect.Milliseconds()
			entry.FirstResponseByteMs = rd.start.Milliseconds()
			entry.TotalMs = rd.total.Milliseconds()

			entry, err = db.InsertHTTPMonitorEntry(entry)

			if err != nil {
				panic(err)
			}
		}
	}
}

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

	monitors, err := db.SelectHTTPMonitors()
	if err != nil {
		log.Fatalln("db:", err)
	}

	for _, m := range monitors {
		switch m.Method {
		case "GET":
			go timeGet(*db, m.Endpoint, m.ID)
		}
	}

	select {}
}
