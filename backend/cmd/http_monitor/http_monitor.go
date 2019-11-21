package main

import (
	"crypto/tls"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/teamwork/reload"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

type requestDuration struct {
	start, connect, dns, tlsHandshake time.Duration
}

func timeGet(url string) {
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

			start = time.Now()

			// Create a new http.Transport to avoid caching results.
			client := &http.Client{Transport: &http.Transport{}}

			start = time.Now()
			r, _ := client.Do(req)
			fmt.Println(url)
			fmt.Printf("Total time: %v\n", time.Since(start))
			fmt.Println(r.StatusCode)
			fmt.Printf("%+v\n", rd)
			fmt.Println("------------------")

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
			go timeGet(m.Endpoint)
		}
	}

	select {}
}
