package core

import (
	"net/url"
	"strings"
)

func ValidateUrl(urlToCheck string) (string, bool) {

	var validatedUrl string

	u, err := url.Parse(urlToCheck)

	if err != nil {
		return urlToCheck, false
	}

	if u.Scheme == "" || u.Host == "" {
		validatedUrl = "http://" + urlToCheck
	} else {
		validatedUrl = urlToCheck
	}

	//remove www. from the url, so that the display of all urls is equivalent
	if strings.Contains(validatedUrl, "www.") {
		validatedUrl = strings.Replace(validatedUrl, "www.", "", 1)
	}
	return validatedUrl, true
}
