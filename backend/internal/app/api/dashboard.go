package api

import (
	"github.com/joaodlf/jsend"
	"net/http"
)

// DashboardTiles returns the information to populate the dashboard tiles.
func (api Resource) DashboardTiles(w http.ResponseWriter, r *http.Request) {
	db := api.AppContext.DB

	httpMonitorsCount, err := db.CountActiveHTTPMonitors()

	if err != nil {
		jsend.Write(w,
			jsend.Data(map[string]interface{}{
				"error": err.Error(),
			}),
			jsend.StatusCode(400),
		)
		return
	}

	jsend.Write(w,
		jsend.StatusCode(200),
		jsend.Data(map[string]interface{}{
			"http_monitors": httpMonitorsCount,
		}),
	)
}
