package validators

// LoginForm represents the /auth/login form details.
type LoginForm struct {
	Username string
	Password string
	Errors   map[string]interface{}
}

// Validate validates the LoginForm data.
func (form *LoginForm) Validate() bool {
	form.Errors = make(map[string]interface{})

	if form.Username == "" {
		form.Errors["username"] = "missing"
	}

	if form.Password == "" {
		form.Errors["password"] = "missing"
	}

	return len(form.Errors) == 0
}
