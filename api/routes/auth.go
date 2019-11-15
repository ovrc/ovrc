package routes

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/joaodlf/jsend"
	"github.com/ovrc/ovrc/validators"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// AuthLogin performs a user login.
func (api Resource) AuthLogin(w http.ResponseWriter, r *http.Request) {
	db := api.AppContext.DB
	form := &validators.LoginForm{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	if form.Validate() == false {
		jsend.Write(w,
			jsend.Data(form.Errors),
			jsend.StatusCode(400),
		)
		return
	}

	user, err := db.SelectUser(form.Username)

	// General error for credential errors, don't want to give too much away (such as incorrect username/password).
	validationError := map[string]interface{}{
		"validation": "Could not validate credentials.",
	}

	if err != nil && err != sql.ErrNoRows {
		jsend.Write(w,
			jsend.Data(validationError),
			jsend.StatusCode(400),
		)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))

	if err != nil {
		jsend.Write(w,
			jsend.Data(validationError),
			jsend.StatusCode(400),
		)
		return
	}

	secure := false
	if api.AppContext.Config.UseSSL == "true" {
		secure = true
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    uuid.New().String(),
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	jsend.Write(w)
}

// AuthLogout effectively logs the user out by "deleting" the session_id cookie.
func (api Resource) AuthLogout(w http.ResponseWriter, r *http.Request) {
	secure := false
	if api.AppContext.Config.UseSSL == "true" {
		secure = true
	}

	// You delete a cookie by setting the expiration to 0.
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Secure:   secure,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)

	jsend.Write(w)
}
