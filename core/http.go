package core

import (
	"net/http"
	"time"

)

func RequestUrl(url string) string {
	timeout := 5 * time.Second

	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)

	if err != nil {
		print(err.Error())
	}

	return resp.Status

	//print(string(resp.StatusCode) + resp.Status)
}
