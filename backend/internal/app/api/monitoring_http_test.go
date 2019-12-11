package api

import (
	"errors"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestMonitoringHTTPDbMock struct {
	model.Datastore
}

func (mdb *TestMonitoringHTTPDbMock) SelectHTTPMonitorEntriesForDashboard() ([]model.HTTPMonitorEntryDashboard, error) {
	mons := []model.HTTPMonitorEntryDashboard{
		{ID: 1, Endpoint: "https://www.google.com/", Method: "GET", AvgTotalMs: 100},
		{ID: 2, Endpoint: "https://www.golang.org/", Method: "GET", AvgTotalMs: 200},
		{ID: 3, Endpoint: "https://www.twitter.com/", Method: "GET", AvgTotalMs: 300},
		{ID: 4, Endpoint: "https://www.github.com/", Method: "GET", AvgTotalMs: 400},
	}

	return mons, nil
}

type TestMonitoringHTTPNoRowsDbMock struct {
	model.Datastore
}

func (mdb *TestMonitoringHTTPNoRowsDbMock) SelectHTTPMonitorEntriesForDashboard() ([]model.HTTPMonitorEntryDashboard, error) {
	var mons []model.HTTPMonitorEntryDashboard
	return mons, errors.New("no rows")
}

// TestMonitoringHTTPSuccess tests for a successful /monitoring/http.
func TestMonitoringHTTPSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := newRequest("GET", "/monitoring/http", nil)

	ac := appcontext.AppContext{DB: &TestMonitoringHTTPDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.MonitoringHTTP).ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, `{"data":{"monitors":[{"avg_total_ms":100,"endpoint":"https://www.google.com/","id":1,"method":"GET"},{"avg_total_ms":200,"endpoint":"https://www.golang.org/","id":2,"method":"GET"},{"avg_total_ms":300,"endpoint":"https://www.twitter.com/","id":3,"method":"GET"},{"avg_total_ms":400,"endpoint":"https://www.github.com/","id":4,"method":"GET"}]},"status":"success"}`, rec.Body.String())
}

// TestMonitoringHTTPSuccess tests for a successful /monitoring/http.
func TestMonitoringHTTPNoRows(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := newRequest("GET", "/monitoring/http", nil)

	ac := appcontext.AppContext{DB: &TestMonitoringHTTPNoRowsDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.MonitoringHTTP).ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
	assert.Equal(t, `{"data":{"error":"no rows"},"status":"fail"}`, rec.Body.String())
}
