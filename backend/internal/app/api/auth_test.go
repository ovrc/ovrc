package api

import (
	"github.com/google/uuid"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type TestAuthLoginSuccessDbMock struct {
	model.Datastore
}

func (mdb *TestAuthLoginSuccessDbMock) SelectUser(username string) (*model.User, error) {
	user := &model.User{
		ID:       1,
		Username: username,
		Password: "$2y$12$tngOXu/YmEXrSactQIDACuiyqL2fj5zohp10ByWPKJRW3tEcqpiPS",
	}
	return user, nil
}

func (mdb *TestAuthLoginSuccessDbMock) UpdateUserSessionID(userID int, sessionID uuid.UUID) error {
	return nil
}

// TestAuthLoginSuccess tests for a successful login.
func TestAuthLoginSuccess(t *testing.T) {
	rec := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "joao")
	data.Set("password", "password")

	req, _ := http.NewRequest("POST", "/auth/login", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ac := appcontext.AppContext{DB: &TestAuthLoginSuccessDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
}

type TestAuthLoginMissingCredentialsDbMock struct {
	model.Datastore
}

// TestAuthLoginMissingCredentials tests for a missing username/password.
func TestAuthLoginMissingCredentials(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/auth/login", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ac := appcontext.AppContext{DB: &TestAuthLoginMissingCredentialsDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
	assert.Equal(t, `{"data":{"password":"missing","username":"missing"},"status":"fail"}`, rec.Body.String())
}

type TestAuthLoginInvalidCredentialsDbMock struct {
	model.Datastore
}

func (mdb *TestAuthLoginInvalidCredentialsDbMock) SelectUser(username string) (*model.User, error) {
	user := &model.User{
		ID:       1,
		Username: username,
		Password: "this-won't-work",
	}
	return user, nil
}

// TestAuthLoginMissingCredentials tests for a missing username/password.
func TestAuthLoginInvalidCredentials(t *testing.T) {
	rec := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "joao")
	data.Set("password", "password")

	req, _ := http.NewRequest("POST", "/auth/login", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ac := appcontext.AppContext{DB: &TestAuthLoginInvalidCredentialsDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
}
