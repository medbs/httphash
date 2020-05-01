package core

import (
	"net/url"
)

func ValidateUrl(urlToCheck string) (string, bool) {

	var validatedUrl string

	u, err := url.Parse(urlToCheck)

	if err != nil {
		return "", false
	}

	// remove the www
	if u.Scheme == "" || u.Host == "" {
		validatedUrl = "http://" + urlToCheck
	} else {
		validatedUrl = urlToCheck
	}
	return validatedUrl, true
}
