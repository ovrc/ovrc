package api

import (
	"github.com/joaodlf/jsend"
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
