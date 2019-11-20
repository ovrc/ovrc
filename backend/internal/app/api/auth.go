package api

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/joaodlf/jsend"
	"github.com/ovrc/ovrc/internal/app/api/validator"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// AuthLogin performs a user login.
func (api Resource) AuthLogin(w http.ResponseWriter, r *http.Request) {
	db := api.AppContext.DB
	form := &validator.LoginForm{
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

	sessionID := uuid.New()

	err = db.UpdateUserSessionID(user.ID, sessionID)

	if err != nil {
		jsend.Write(w,
			jsend.StatusCode(500),
			jsend.Message(err.Error()),
		)
		return
	}

	secure := false
	if api.AppContext.Config.UseSSL == "true" {
		secure = true
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID.String(),
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	jsend.Write(w, jsend.Data(map[string]interface{}{
		"username": user.Username,
	}))
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
