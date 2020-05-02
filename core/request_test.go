package core

import "testing"

func TestHashRequestResponse(t *testing.T) {
	response := HashRequestResponse("request response")

	if response == "" {
		t.Errorf("hashing the request response should not return an empty string")
	}
}

func TestRequestUrl(t *testing.T) {
	//first level of validation, same as ValidateUrl() function
	_, _, s1 := RequestUrl("www.twitter.com")
	_, _, s2 := RequestUrl("google.com")
	_, _, s3 := RequestUrl("http://amazon.com")
	_, _, s4 := RequestUrl("http://www.yahoo.com")

	//second level of validation
	_, _, s5 := RequestUrl("htttp://www.yahoo.com")

	_, _, s6 := RequestUrl("http://www.yahoo//com")

	if !s1 {
		t.Errorf("error when sending request to url with the format www.url.com")
	}

	if !s2 {
		t.Errorf("error when sending request to url with the format url.com")
	}

	if !s3 {
		t.Errorf("error when sending request to url with the format http://url.com")
	}

	if !s4 {
		t.Errorf("error when sending request to url with the format http://www.url.com")
	}

	if s5 {
		t.Errorf("error when sending request to url with the format htttp://www.url/com")
	}

	if s6 {
		t.Errorf("error when sending request to url with the format http://www.url//com")
	}
}
