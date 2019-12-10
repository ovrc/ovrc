package api

import (
	"context"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ovrc/ovrc/internal/model"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// newRequest creates a new http.request with the correct Content-Type header.
func newRequest(method, url string, body url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}

// newSessionRequest creates a new http.request with a user context to mimic a logged in user.
func newSessionRequest(method, url string, body url.Values) (*http.Request, context.Context) {
	req, _ := newRequest(method, url, body)
	ctx := context.WithValue(req.Context(), "user", &model.User{
		ID:        1,
		DtCreated: time.Time{},
		DtUpdated: pq.NullTime{},
		Username:  "joao",
		Password:  "$2y$12$tngOXu/YmEXrSactQIDACuiyqL2fj5zohp10ByWPKJRW3tEcqpiPS",
		SessionID: uuid.UUID{},
	})

	return req, ctx
}
