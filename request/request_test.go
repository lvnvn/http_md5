package request

import (
	"errors"
	"regexp"
	"testing"
)

const testUrl = "http://baroquemusiclibrary.com"
const testUrlWithoutScheme = "baroquemusiclibrary.com"

func TestNormalizeUrl(t *testing.T) {
	result, _ := NormalizeUrl(testUrl)
	expected := testUrl
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestNormalizeUrlWithoutScheme(t *testing.T) {
	result, _ := NormalizeUrl(testUrlWithoutScheme)
	expected := testUrl
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestNormalizeUrlInvalid(t *testing.T) {
	_, err := NormalizeUrl("a|b")
	expected := "parse \"http://a|b\": invalid character \"|\" in host name"
	if err.Error() != expected {
		t.Errorf("got %s, want %s", err.Error(), expected)
	}
}

type MockClient struct {
	ResponseBody  []byte
	ReturnedError error
}

func (c MockClient) MakeRequest(url string) ([]byte, error) {
	return c.ResponseBody, c.ReturnedError
}

func TestHttpToMD5Positive(t *testing.T) {
	responseBody := []byte(`
		<!doctype html>
		<html lang="en">
		  <head><title>Test</title></head>
		  <body>test test</body>
		</html>
	`)
	client := MockClient{
		ResponseBody:  responseBody,
		ReturnedError: nil,
	}
	result, err := HttpToMD5(testUrlWithoutScheme, client)
	if err != nil {
		t.Errorf("expected no error")
	}

	expected := testUrl
	if result.Url != expected {
		t.Errorf("got %s, want %s", result.Url, expected)
	}

	matched, _ := regexp.MatchString("[a-f0-9]{32}", result.Hash)
	if matched != true {
		t.Errorf("%s is not an MD5 hash", result.Hash)
	}

	result2, _ := HttpToMD5(testUrl, client)
	if result != result2 {
		t.Errorf("hashes for two requests don't match")
	}
}

func TestHttpToMD5RequestError(t *testing.T) {
	errMessage := "Test error message"
	client := MockClient{
		ResponseBody:  make([]byte, 0),
		ReturnedError: errors.New(errMessage),
	}
	result, err := HttpToMD5(testUrlWithoutScheme, client)
	if err.Error() != errMessage {
		t.Errorf("expected error %s, got %s", errMessage, err.Error())
	}
	if result.Url != testUrl {
		t.Errorf("got %s, want %s", result.Url, testUrl)
	}
}
