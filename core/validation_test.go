package core

import "testing"

func TestValidateUrl(t *testing.T) {
	_, validUrl1 := ValidateUrl("www.twitter.com")
	_, validUrl2 := ValidateUrl("google.com")
	_, validUrl3 := ValidateUrl("http://amazon.com")
	_, validUrl4 := ValidateUrl("http://www.yahoo.com")

	if !validUrl1 {
		t.Errorf("error when validating url format www.url.com")
	}

	if !validUrl2 {
		t.Errorf("error when validating url format url.com")
	}

	if !validUrl3 {
		t.Errorf("error when validating url format http://url.com")
	}

	if !validUrl4 {
		t.Errorf("error when validating url format http://www.url.com")
	}

}
