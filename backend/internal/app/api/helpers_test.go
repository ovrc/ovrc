package api

import (
	"net/http"
	"net/url"
	"strings"
)

// newRequest creates a new http.request with the correct Content-Type header.
func newRequest(method, url string, body url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}
