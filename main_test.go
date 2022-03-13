package main

import (
	"io/ioutil"
	"testing"
)

type testHttpClient struct {
}

func (client testHttpClient) get() string {
	content, err := ioutil.ReadFile("test-asset/google-homepage.html")

	if err != nil {
	}

	// Convert []byte to string and print to screen
	return string(content)
}

func TestScriptFound(t *testing.T) {
	location := getWebPageFindStringAndReturnLocation(testHttpClient{}, "script")

	if location != 273 {
		t.Fatal("failed")
	}
}
