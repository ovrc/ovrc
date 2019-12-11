package api

import (
	"github.com/joaodlf/jsend"
	"github.com/ovrc/ovrc/internal/app/api/validator"
	"github.com/ovrc/ovrc/internal/model"
	"net/http"
)

// MonitoringHTTP returns a list of http monitors with related figures for the dashboard.
func (api Resource) MonitoringHTTP(w http.ResponseWriter, r *http.Request) {
	entries, err := api.AppContext.DB.SelectHTTPMonitorEntriesForDashboard()

	if err != nil {
		jsend.Write(w,
			jsend.Data(map[string]interface{}{
				"error": err.Error(),
			}),
			jsend.StatusCode(400),
		)
		return
	}

	var entryList []map[string]interface{}
	for _, row := range entries {
		entryList = append(entryList, map[string]interface{}{
			"id":           row.ID,
			"endpoint":     row.Endpoint,
			"method":       row.Method,
			"avg_total_ms": row.AvgTotalMs,
		})
	}

	jsend.Write(w,
		jsend.StatusCode(200),
		jsend.Data(map[string]interface{}{
			"monitors": entryList,
		}),
	)
}

// MonitoringHTTPAdd adds a new HTTP monitor. //TODO: tests
func (api Resource) MonitoringHTTPAdd(w http.ResponseWriter, r *http.Request) {
	form := &validator.HTTPMonitorAdd{
		Method: r.FormValue("method"),
		URL:    r.FormValue("url"),
	}

	if form.Validate() == false {
		jsend.Write(w,
			jsend.Data(form.Errors),
			jsend.StatusCode(400),
		)
		return
	}

	mon := model.HTTPMonitor{
		Method:   form.Method,
		Endpoint: form.URL,
	}

	httpMonitor, err := api.AppContext.DB.InsertHTTPMonitor(mon)

	if err != nil {
		jsend.Write(w,
			jsend.Message(err.Error()),
			jsend.StatusCode(500),
		)
		return
	}

	jsend.Write(w,
		jsend.StatusCode(200),
		jsend.Data(map[string]interface{}{
			"id": httpMonitor.ID,
		}),
	)
}
