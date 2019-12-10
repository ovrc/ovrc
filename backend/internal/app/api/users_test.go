package api

import (
	"errors"
	"github.com/ovrc/ovrc/internal/appcontext"
	"github.com/ovrc/ovrc/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestUsersSuccessDbMock struct {
	model.Datastore
}

func (mdb *TestUsersSuccessDbMock) SelectUsersForAdmin() ([]model.User, error) {
	dtCreated, _ := time.Parse("2006-01-02 15:04:05", "2019-01-02 03:00:00")

	users := []model.User{
		{ID: 1, Username: "joao", DtCreated: dtCreated},
		{ID: 2, Username: "notjoao", DtCreated: dtCreated},
		{ID: 3, Username: "admin", DtCreated: dtCreated},
	}

	return users, nil
}

type TestUsersFailDbMock struct {
	model.Datastore
}

func (mdb *TestUsersFailDbMock) SelectUsersForAdmin() ([]model.User, error) {
	return []model.User{}, errors.New("fail")
}

func TestUsersMeSuccess(t *testing.T) {
	rec := httptest.NewRecorder()
	req, ctx := newSessionRequest("GET", "/users/me", nil)

	ac := appcontext.AppContext{}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.UsersMe).ServeHTTP(rec, req.WithContext(ctx))

	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, `{"data":{"username":"joao"},"status":"success"}`, rec.Body.String())
}

func TestUsersSuccess(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := newRequest("GET", "/users", nil)

	ac := appcontext.AppContext{DB: &TestUsersSuccessDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.Users).ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, `{"data":{"users":[{"dt_created":"2019-01-02 03:00:00","id":1,"username":"joao"},{"dt_created":"2019-01-02 03:00:00","id":2,"username":"notjoao"},{"dt_created":"2019-01-02 03:00:00","id":3,"username":"admin"}]},"status":"success"}`, rec.Body.String())
}

func TestUsersFail(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := newRequest("GET", "/users", nil)

	ac := appcontext.AppContext{DB: &TestUsersFailDbMock{}}
	apiResource := Resource{AppContext: ac}

	http.HandlerFunc(apiResource.Users).ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
}
