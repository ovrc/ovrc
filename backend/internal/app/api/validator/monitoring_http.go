package validator

// LoginForm represents the /auth/login form details.
type HTTPMonitorAdd struct {
	Method string
	URL    string
	Errors map[string]interface{}
}

// Validate validates the LoginForm data.
func (form *HTTPMonitorAdd) Validate() bool {
	form.Errors = make(map[string]interface{})

	if form.Method == "" {
		form.Errors["method"] = "missing"
	}

	validMethod := false
	switch form.Method {
	case "GET", "POST":
		validMethod = true
	}

	if !validMethod {
		form.Errors["method"] = "invalid value"
	}

	if form.URL == "" {
		form.Errors["url"] = "missing"
	}

	return len(form.Errors) == 0
}
