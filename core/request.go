package core

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"
)

func HashRequestResponse(requestResponse string) string {

	hash := md5.Sum([]byte(requestResponse))
	return hex.EncodeToString(hash[:])

}

func RequestUrl(url string) (string, string, bool) {
	timeout := 5 * time.Second

	client := http.Client{
		Timeout: timeout,
	}

	validatedUrl, isValid := ValidateUrl(url)

	if !isValid {
		return "", "", false
	}

	resp, err := client.Get(validatedUrl)

	if err != nil {
		//print(err.Error())
		return "", err.Error(), false
	}

	return validatedUrl, resp.Status, true

	//print(string(resp.StatusCode) + resp.Status)
}
