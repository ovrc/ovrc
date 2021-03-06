package api

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
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

type TestAuthLoginMissingCredentialsDbMock struct {
	model.Datastore
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

type TestAuthLoginInvalidUsernameDbMock struct {
	model.Datastore
}

func (mdb *TestAuthLoginInvalidUsernameDbMock) SelectUser(username string) (*model.User, error) {
	user := &model.User{}
	return user, errors.New("invalid user")
}

type TestAuthLoginFailedUpdateDbMock struct {
	model.Datastore
}

func (mdb *TestAuthLoginFailedUpdateDbMock) SelectUser(username string) (*model.User, error) {
	user := &model.User{
		ID:        1,
		Username:  username,
		Password:  "$2y$12$tngOXu/YmEXrSactQIDACuiyqL2fj5zohp10ByWPKJRW3tEcqpiPS",
		SessionID: uuid.New(),
	}
	return user, nil
}

func (mdb *TestAuthLoginFailedUpdateDbMock) UpdateUserSessionID(userID int, sessionID uuid.UUID) error {
	return errors.New("failed to update")
}

// TestAuthLoginSuccess tests for a successful login.
func TestAuthLoginSuccess(t *testing.T) {
	rec := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "joao")
	data.Set("password", "password")

	req, _ := newRequest("POST", "/auth/login", data)

	ac := appcontext.AppContext{DB: &TestAuthLoginSuccessDbMock{}}
	ac.Config.UseSSL = "true"
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
}

// TestAuthLoginMissingCredentials tests for a missing username/password.
func TestAuthLoginMissingCredentials(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := newRequest("POST", "/auth/login", nil)

	ac := appcontext.AppContext{DB: &TestAuthLoginMissingCredentialsDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
	assert.Equal(t, `{"data":{"password":"missing","username":"missing"},"status":"fail"}`, rec.Body.String())
}

// TestAuthLoginInvalidCredentials tests for an invalid username/password combination.
func TestAuthLoginInvalidCredentials(t *testing.T) {
	rec := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "joao")
	data.Set("password", "password")

	req, _ := newRequest("POST", "/auth/login", data)

	ac := appcontext.AppContext{DB: &TestAuthLoginInvalidCredentialsDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
	assert.Equal(t, `{"data":{"validation":"Could not validate credentials."},"status":"fail"}`, rec.Body.String())
}

// TestAuthLoginInvalidUsername tests for an invalid username.
func TestAuthLoginInvalidUsername(t *testing.T) {
	rec := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "joao")
	data.Set("password", "password")

	req, _ := newRequest("POST", "/auth/login", data)

	ac := appcontext.AppContext{DB: &TestAuthLoginInvalidUsernameDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
	assert.Equal(t, `{"data":{"validation":"Could not validate credentials."},"status":"fail"}`, rec.Body.String())
}

// TestAuthLoginFailedUpdate tests for a failing user update.
func TestAuthLoginFailedUpdate(t *testing.T) {
	rec := httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "joao")
	data.Set("password", "password")

	req, _ := newRequest("POST", "/auth/login", data)

	ac := appcontext.AppContext{DB: &TestAuthLoginFailedUpdateDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogin).ServeHTTP(rec, req)

	assert.Equal(t, 500, rec.Code)
}

func TestAuthLogout(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := newRequest("GET", "/auth/logout", nil)

	ac := appcontext.AppContext{}
	ac.Config.UseSSL = "true"
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.AuthLogout).ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
}
