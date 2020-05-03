package core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func SendRequestHashResponse(url string) {
	validUrl, requestResponse, isValid := RequestUrl(url)
	if !isValid {
		fmt.Print(validUrl + " Url format is wrong or cannot be requested \n")
	} else {
		hash := HashRequestResponse(requestResponse)

		fmt.Println(validUrl + " " + " " + hash)

	}
}

//RequestUrl sends request to a given url, returns the url, the response of the request, and a boolean indicating if the request is successful or not
func RequestUrl(url string) (string, string, bool) {
	timeout := 5 * time.Second

	client := http.Client{
		Timeout: timeout,
	}

	validatedUrl, isValid := ValidateUrl(url)

	if !isValid {
		return validatedUrl, "", false
	}

	resp, err := client.Get(validatedUrl)

	//always return validatedUrl even it has wrong format, to be displayed to the user
	if err != nil {
		return validatedUrl, err.Error(), false
	}

	//close reading body at the end of the function to prevent resources leaks
	defer resp.Body.Close()

	//read the body of the response
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return validatedUrl, err.Error(), false
	}

	return validatedUrl, string(body), true

}

// HashRequestResponse takes a string argument and returns the md5 hash of it
func HashRequestResponse(requestResponse string) string {

	hash := md5.Sum([]byte(requestResponse))
	return hex.EncodeToString(hash[:])

}
