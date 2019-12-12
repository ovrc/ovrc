package validator

// HTTPMonitorAdd represents the POST /monitor/http form details.
type HTTPMonitorAdd struct {
	Method string
	URL    string
	Errors map[string]interface{}
}

// HTTPMonitorList represents the GET /monitor/http form details.
type HTTPMonitorList struct {
	Period string
	Errors map[string]interface{}
}

// Validate validates the HTTPMonitorAdd data.
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

// Validate validates the HTTPMonitorList data.
func (form *HTTPMonitorList) Validate() bool {
	form.Errors = make(map[string]interface{})

	validPeriod := false
	switch form.Period {
	case "hour1", "hour3", "hour6", "hour12", "hour24":
		validPeriod = true
		break
	}

	if !validPeriod {
		form.Errors["period"] = "invalid value"
	}

	return len(form.Errors) == 0
}
