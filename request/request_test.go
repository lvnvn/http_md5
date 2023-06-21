package request

import (
	"fmt"
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

func TestHttpToMD5(t *testing.T) {
	// TODO: mock client.Get for test stability
	result, err := HttpToMD5(testUrlWithoutScheme)
	if err != nil {
		fmt.Printf("%s",err)
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

	result2, _ := HttpToMD5(testUrl)
	if result != result2 {
		t.Errorf("hashes for two requests don't match")
	}
}
