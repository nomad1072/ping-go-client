package main

import (
	"os"
	"testing"
)

func TestArgs(t *testing.T) {
	expectedHostname := "amazon.com"
	expectedCount := 50
	os.Args = append(os.Args, "-hostname=amazon.com")
	os.Args = append(os.Args, "-count=50")

	actualHostname, actualCount, _, _ := fetchArgs()
	if actualHostname != expectedHostname {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expectedHostname, actualHostname)
	}
	if expectedCount != actualCount {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expectedCount, actualCount)
	}
}
