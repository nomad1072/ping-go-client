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

	t.Run("testCommandLineArgs", func(t *testing.T) {
		actualHostname, actualCount, _, _ := fetchArgs()
		if actualHostname != expectedHostname {
			t.Errorf("Test failed, expected: '%s', got: '%s'", expectedHostname, actualHostname)
		}
		if expectedCount != actualCount {
			t.Errorf("Test failed, expected: '%d', got: '%d'", expectedCount, actualCount)
		}
	})

	t.Run("validateCommandLineArgs", func(t *testing.T) {
		message, _ := validateArgs(-1, "googkle.com", 34)
		wrongHostName := "invalid IP address/hostname provided"
		if message != wrongHostName {
			t.Errorf("Test failed, provided incorrect hostname but received ambiguous output")
		}
	})

	t.Run("NegativeIPInput=-27838838383", func(t *testing.T) {
		message, _ := validateArgs(-1, "-27838838383", 34)
		wrongHostName := "invalid IP address/hostname provided"
		if message != wrongHostName {
			t.Errorf("Test failed, provided incorrect hostname but received ambiguous output")
		}
	})

	t.Run("validateCommandLineArgs", func(t *testing.T) {
		message, _ := validateArgs(-1, "google.com", 34)
		wrongHostName := "count accepts positive integers"
		if message != wrongHostName {
			t.Errorf("Test failed, provided incorrect count value but client accepts given value")
		}
	})

	t.Run("InvalidTTLValue=-38", func(t *testing.T) {
		message, _ := validateArgs(-1, "-27838838383", -34)
		wrongHostName := "invalid value specified for ttl"
		if message != wrongHostName {
			t.Errorf("Test failed, provided incorrect hostname but received ambiguous output")
		}
	})

	t.Run("InvalidTTLValue=289", func(t *testing.T) {
		message, _ := validateArgs(-1, "-27838838383", 289)
		wrongHostName := "invalid value specified for ttl"
		if message != wrongHostName {
			t.Errorf("Test failed, provided incorrect hostname but received ambiguous output")
		}
	})
}
