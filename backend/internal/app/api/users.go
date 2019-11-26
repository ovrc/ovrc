package api

import (
	"github.com/joaodlf/jsend"
	"github.com/ovrc/ovrc/internal/model"
	"net/http"
)

// UsersMe returns the logged in user details.
func (api Resource) UsersMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)

	jsend.Write(w,
		jsend.StatusCode(200),
		jsend.Data(map[string]interface{}{
			"username": user.Username,
		}),
	)
}

func (api Resource) Users(w http.ResponseWriter, r *http.Request) {
	users, err := api.AppContext.DB.SelectUsersForAdmin()

	if err != nil {
		jsend.Write(w,
			jsend.Data(map[string]interface{}{
				"error": err.Error(),
			}),
			jsend.StatusCode(400),
		)
		return
	}

	var userList []map[string]interface{}
	for _, row := range users {
		userList = append(userList, map[string]interface{}{
			"id":         row.ID,
			"username":   row.Username,
			"dt_created": row.DtCreated.Format("2006-01-02 03:04:05"),
		})
	}

	jsend.Write(w,
		jsend.StatusCode(200),
		jsend.Data(map[string]interface{}{
			"users": userList,
		}),
	)
}
