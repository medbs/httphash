package core

import (
	"net/url"
	"strings"
)

func ValidateUrl(url string) (string, bool) {

	/* possible format :
	www.url.com issue
	url.com issue
	http://url.com  ok
	http://www.url.com ok */

	if ! isValidUrl(url) {

		if !strings.Contains(url, "http") {
			return "http://" + url, true
		}
	}
	return " ", false

}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
