package api

import (
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersMeSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	req, ctx := newSessionRequest("GET", "/users/me", nil)

	ac := appcontext.AppContext{}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.UsersMe).ServeHTTP(rec, req.WithContext(ctx))

	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, `{"data":{"username":"joao"},"status":"success"}`, rec.Body.String())
}
