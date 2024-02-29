package api_test

import (
	"testing"

	"frontendmasters.com/courses/go-basics/cryptomasters/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
